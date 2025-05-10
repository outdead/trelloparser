package trello

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/outdead/trelloparser/libs/trello/entity"
)

type Config struct{}

type Parser struct {
	config Config
}

func NewParser(scg Config) *Parser {
	return &Parser{config: scg}
}

// Parse reads and parses a Trello board JSON file into a Dashboard struct.
// It takes the board filename as input and returns a pointer to the Dashboard and an error.
// The function performs the following steps:
//  1. Reads the JSON file from disk
//  2. Unmarshals the JSON into a Dashboard struct
//  3. Sorts the cards by due date in descending order (newest first)
//
// Note: Cards cannot be sorted by creation date as Trello doesn't export complete creation records.
// Returns:
//   - *entity.Dashboard: Pointer to the populated dashboard structure
//   - error: Any error that occurred during file reading or JSON parsing.
//
// TODO: Add sortCards param.
func (m *Parser) Parse(boardName string) (*entity.Dashboard, error) {
	file, err := os.ReadFile(boardName)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", err, boardName)
	}

	var dash entity.Dashboard
	if err := json.Unmarshal(file, &dash); err != nil {
		return nil, err
	}

	// Sort cards by due date (newest first).
	// There is no way to sort by card creation date. Some creation date records
	// can be retrieved from the Actions slice by type "createMarkdownCard" and IDCard,
	// but Trello does not export all the corresponding records to Actions.
	sort.Slice(dash.Cards, func(i, j int) bool {
		return dash.Cards[i].Due.After(dash.Cards[j].Due)
	})

	return &dash, nil
}

// ParseAndAggregate combines the parsing and aggregation steps into a single operation.
// It first parses the raw data for the given board name, then aggregates the parsed data
// into a structured dashboard format with proper relationships between entities.
//
// Parameters:
//   - boardName: Name of the board to parse and aggregate
//
// Returns:
//   - *entity.Dashboard: Fully populated dashboard structure with all relationships
//   - error: Any error that occurred during parsing.
func (m *Parser) ParseAndAggregate(boardName string) (*entity.Dashboard, error) {
	dash, err := m.Parse(boardName)
	if err != nil {
		return nil, fmt.Errorf("failed to parse dashboard %s: %w", boardName, err)
	}

	return m.Aggregate(dash), nil
}

// Aggregate processes a parsed dashboard structure to establish relationships between
// entities. It:
//   - Creates lookup caches for efficient access
//   - Attaches checklists to their respective cards
//   - Groups cards under their parent lists
//   - Filters out closed items
//
// Parameters:
//   - dash: Pointer to the initially parsed dashboard structure
//
// Returns:
//   - *entity.Dashboard: The same dashboard with all relationships established.
func (m *Parser) Aggregate(dash *entity.Dashboard) *entity.Dashboard {
	// Create a map of checklists by their ID for efficient lookup.
	checklistsCache := make(map[string]entity.Checklist)
	for i := range dash.Checklists {
		checklistsCache[dash.Checklists[i].ID] = dash.Checklists[i]
	}

	// Create a map of non-closed lists by their ID.
	listsCache := make(map[string]entity.List, len(dash.Lists))

	for i := range dash.Lists {
		list := dash.Lists[i]

		if list.Closed {
			continue
		}

		listsCache[list.ID] = list
	}

	// Process all cards to:
	// 1. Attach their checklists
	// 2. Group them under their parent lists.
	for i := range dash.Cards {
		card := dash.Cards[i]

		if card.Closed {
			continue
		}

		// Attach all checklists to the card.
		for _, checkID := range card.IDChecklists {
			checklist, ok := checklistsCache[checkID]
			if !ok {
				continue
			}

			card.Checklists = append(card.Checklists, checklist)
		}

		// Find the card's parent list.
		cachedList, ok := listsCache[card.IDList]
		if !ok {
			continue
		}

		// Add the card to its parent list in the cache.
		if cachedList.ID == card.IDList {
			cachedList.Cards = append(cachedList.Cards, card)
			listsCache[card.IDList] = cachedList
		}
	}

	// Update the original lists with the aggregated card data.
	for i := range dash.Lists {
		list := &dash.Lists[i]

		cachedList, ok := listsCache[list.ID]
		if !ok {
			continue
		}

		list.Cards = append(list.Cards, cachedList.Cards...)
	}

	return dash
}
