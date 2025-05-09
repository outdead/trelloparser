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

// Markdown struct holds configuration and logger for markdown generation.
type Markdown struct {
	config *config.Config
	logger *logger.Logger
}

// New creates and returns a new Markdown instance with the given config and logger.
func New(config *config.Config, logger *logger.Logger) *Markdown {
	m := Markdown{
		config: config,
		logger: logger,
	}

	return &m
}

// CreateMarkdown generates markdown files from a Trello board JSON file.
func (m *Markdown) CreateMarkdown(boardName string) error {
	file, err := os.ReadFile(boardName)
	if err != nil {
		return err
	}

	var dash entity.Dashboard
	if err = json.Unmarshal(file, &dash); err != nil {
		return err
	}

	// Sort cards by due date (newest first).
	// There is no way to sort by card creation date. Some creation date records
	// can be retrieved from the Actions slice by type "createCard" and IDCard,
	// but Trello does not export all the corresponding records to Actions.
	sort.Slice(dash.Cards, func(i, j int) bool {
		return dash.Cards[i].Due.After(dash.Cards[j].Due)
	})

	// Create a map of non-closed cards by their ID.
	cards := make(map[string]entity.Card)
	for _, card := range dash.Cards {
		if !card.Closed {
			cards[card.ID] = card
		}
	}

	// Create a map of checklists by their ID.
	checklists := make(map[string]entity.Checklist)
	for _, check := range dash.Checklists {
		checklists[check.ID] = check
	}

	// Create main markdown file for the board.

	// Create output file in the configured directory.
	// TODO: move directory name to config or application flags.
	boardFile, err := os.Create(m.config.App.HomeDirectory + "/.tmp/data/response/" + dash.Name + ".md")
	if err != nil {
		return err
	}
	defer boardFile.Close()

	boardMarkdown := markdown.NewMarkdown(boardFile)

	for _, list := range dash.Lists {
		if list.Closed {
			continue
		}

		boardMarkdown.H2(list.Name)

		// Process each card in the list.
		// TODO: This is a bad code. We process each card extra len(dash.Lists) times.
		for _, card := range dash.Cards {
			if card.IDList == list.ID {
				title := "[[" + card.Name + "]]"

				if !card.Due.IsZero() {
					title += " @{" + card.Due.Format("2006-01-02") + "}"
				}

				// Add card as checkbox item (checked if completed).
				boardMarkdown.CheckBox([]markdown.CheckBoxSet{
					{Checked: card.DueComplete, Text: title},
				})
			}
		}

		// Add empty line after each list for pretty markdown file viewing.
		boardMarkdown.PlainText("")
	}

	if err := boardMarkdown.Build(); err != nil {
		return err
	}

	// Create individual markdown files for each checklist.
	// TODO: revert logic: process cards first and associate checklists to the card.
	for _, checklist := range dash.Checklists {
		// Get the card associated with this checklist.
		card, ok := cards[checklist.IDCard]
		if !ok || card.Closed {
			continue
		}

		cardFile, err := os.Create(m.config.App.HomeDirectory + "/.tmp/data/response/tododata/" + card.Name + ".md")
		if err != nil {
			return err
		}

		cardMarkdown := markdown.NewMarkdown(cardFile)

		cardMarkdown.H1(card.Name)

		for _, item := range checklist.CheckItems {
			cardMarkdown.CheckBox([]markdown.CheckBoxSet{
				{Checked: item.State == "complete", Text: item.Name},
			})
		}

		date := ""
		if !card.Due.IsZero() {
			date = card.Due.Format("2006-01-02")
		}

		// Add footer with date and tags.
		// TODO: Add a switch to the configuration file.
		cardMarkdown.PlainText("").PlainText(fmt.Sprintf(cardFooterTemplate, date))

		if err := cardMarkdown.Build(); err != nil {
			_ = cardFile.Close()

			return err
		}

		if err := cardFile.Close(); err != nil {
			return err
		}
	}

	return nil
}
