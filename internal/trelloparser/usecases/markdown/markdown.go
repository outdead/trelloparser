package markdown

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/nao1215/markdown"
	"github.com/outdead/trelloparser/internal/trelloparser/config"
	"github.com/outdead/trelloparser/internal/trelloparser/entity"
	"github.com/outdead/trelloparser/internal/utils/logger"
)

// cardFooterTemplate defines the template for the card footer in markdown format.
const cardFooterTemplate = `---

- date: %s
- tags: #todo`

// ChecklistsCache is a map that stores checklists by their ID for quick lookup.
// The map key is the checklist ID (string) and the value is the corresponding Checklist entity.
// This cache helps avoid repeated linear searches when processing multiple cards.
type ChecklistsCache map[string]entity.Checklist

// Markdown struct holds configuration and logger for markdown generation.
type Markdown struct {
	config *config.Config
	logger *logger.Logger
}

// New creates and returns a new Markdown instance with the given config and logger.
func New(cfg *config.Config, log *logger.Logger) *Markdown {
	m := Markdown{
		config: cfg,
		logger: log,
	}

	return &m
}

// CreateMarkdown generates markdown files from a Trello board JSON file.
func (m *Markdown) CreateMarkdown(boardName string) error {
	dash, err := m.parseDashboardJSON(boardName)
	if err != nil {
		return err
	}

	// Create a map of checklists by their ID.
	checklistsCache := make(ChecklistsCache)
	for i := range dash.Checklists {
		checklistsCache[dash.Checklists[i].ID] = dash.Checklists[i]
	}

	return m.createMarkdownFiles(dash, checklistsCache)
}

func (m *Markdown) createMarkdownFiles(dash *entity.Dashboard, checklistsCache ChecklistsCache) error {
	// Create markdown file for the Dashboard.
	// TODO: move directory name to config or application flags.
	boardFile, err := os.Create(m.config.App.HomeDirectory + "/.tmp/data/response/" + dash.Name + ".md")
	if err != nil {
		return err
	}
	defer boardFile.Close()

	boardMarkdown := markdown.NewMarkdown(boardFile)

	for i := range dash.Lists {
		list := &dash.Lists[i]

		if list.Closed {
			continue
		}

		boardMarkdown.H2(list.Name)

		// Process each card in the list.
		// TODO: We process each card extra len(dash.Lists) times.
		// Good place for optimization with aggregated lists or/and cards.
		for j := range dash.Cards {
			card := dash.Cards[j]

			if card.IDList != list.ID {
				continue
			}

			title := "[[" + card.Name + "]]"

			if !card.Due.IsZero() {
				title += " @{" + card.Due.Format("2006-01-02") + "}"
			}

			// Add card as checkbox item (checked if completed).
			boardMarkdown.CheckBox([]markdown.CheckBoxSet{
				{Checked: card.DueComplete, Text: title},
			})

			cardName := m.config.App.HomeDirectory + "/.tmp/data/response/tododata/" + card.Name + ".md"

			if err := m.createMarkdownCard(cardName, &card, checklistsCache); err != nil {
				return err
			}
		}

		// Add empty line after each list for pretty markdown file viewing.
		boardMarkdown.PlainText("")
	}

	return boardMarkdown.Build()
}

// createMarkdownCard create individual markdown files received Card.
func (m *Markdown) createMarkdownCard(fileName string, card *entity.Card, checklistsCache ChecklistsCache) error {
	if card.Closed {
		return nil
	}

	cardFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer cardFile.Close()

	cardMarkdown := markdown.NewMarkdown(cardFile)

	cardMarkdown.H1(card.Name)

	if card.Desc != "" {
		cardMarkdown.PlainText(card.Desc)
	}

	cardMarkdown.PlainText("")

	for _, checkID := range card.IDChecklists {
		checklist, ok := checklistsCache[checkID]
		if !ok {
			continue
		}

		cardMarkdown.H2(checklist.Name)

		for i := range checklist.CheckItems {
			if checklist.CheckItems[i].IDChecklist == checkID {
				cardMarkdown.CheckBox([]markdown.CheckBoxSet{
					{Checked: checklist.CheckItems[i].State == "complete", Text: checklist.CheckItems[i].Name},
				})
			}
		}

		cardMarkdown.PlainText("")
	}

	date := ""
	if !card.Due.IsZero() {
		date = card.Due.Format("2006-01-02")
	}

	// Add footer with date and tags.
	// TODO: Add a switch to the configuration file.
	cardMarkdown.PlainText(fmt.Sprintf(cardFooterTemplate, date))

	if err := cardMarkdown.Build(); err != nil {
		return err
	}

	return nil
}

// parseDashboardJSON reads and parses a Trello board JSON file into a Dashboard struct.
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
func (m *Markdown) parseDashboardJSON(boardName string) (*entity.Dashboard, error) {
	file, err := os.ReadFile(boardName)
	if err != nil {
		return nil, err
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
