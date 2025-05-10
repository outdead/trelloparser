package markdown

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/nao1215/markdown"
	"github.com/outdead/golibs/files"
	"github.com/outdead/trelloparser/internal/trelloparser/config"
	"github.com/outdead/trelloparser/internal/utils/logger"
	"github.com/outdead/trelloparser/libs/trello"
	"github.com/outdead/trelloparser/libs/trello/entity"
)

var ErrNotExistResultFolder = errors.New("not exist result folder")

// Markdown struct holds configuration and logger for markdown generation.
type Markdown struct {
	config *config.Config
	logger *logger.Logger
	parser *trello.Parser
}

// New creates and returns a new Markdown instance with the given config and logger.
func New(cfg *config.Config, log *logger.Logger) *Markdown {
	m := Markdown{
		config: cfg,
		logger: log,
		parser: trello.NewParser(cfg.Trello),
	}

	_ = m.prepareDataDirectory()

	return &m
}

// CreateMarkdown generates markdown files from a Trello board JSON file.
func (m *Markdown) CreateMarkdown(boardName string) error {
	dash, err := m.parser.ParseAndAggregate(boardName)
	if err != nil {
		return err
	}

	return m.createMarkdownFiles(dash)
}

// createMarkdownFiles generates markdown documentation for a Trello dashboard.
// It creates:
//   - A main dashboard file with all lists and cards
//   - Individual card files in a subdirectory
//
// The output format can be customized based on configuration (standard markdown or Obsidian-flavored).
//
// Parameters:
//   - dash: Pointer to the dashboard entity containing all data to document
//
// Returns:
//   - error: Returns any file creation or writing errors encountered.
func (m *Markdown) createMarkdownFiles(dash *entity.Dashboard) error {
	// Create markdown file for the Dashboard.
	boardFile, err := os.Create(m.config.Markdown.ResultDirectory + "/" + dash.Name + ".md")
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
		for j := range list.Cards {
			card := list.Cards[j]

			title := m.getCardTitle(&card)

			// Add card as checkbox item (checked if completed).
			boardMarkdown.CheckBox([]markdown.CheckBoxSet{
				{Checked: card.DueComplete, Text: title},
			})

			cardName := m.config.Markdown.ResultDirectory + "/cards/" + card.Name + ".md"

			if err := m.createMarkdownCard(cardName, &card); err != nil {
				return err
			}
		}

		// Add empty line after each list for pretty markdown file viewing.
		boardMarkdown.PlainText("")
	}

	return boardMarkdown.Build()
}

// createMarkdownCard create individual markdown files received Card.
func (m *Markdown) createMarkdownCard(fileName string, card *entity.Card) error {
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

	for i := range card.Checklists {
		checklist := &card.Checklists[i]

		cardMarkdown.H2(checklist.Name)

		for i := range checklist.CheckItems {
			cardMarkdown.CheckBox([]markdown.CheckBoxSet{
				{Checked: checklist.CheckItems[i].State == "complete", Text: checklist.CheckItems[i].Name},
			})
		}

		cardMarkdown.PlainText("")
	}

	m.addCardFooter(cardMarkdown, card)

	return cardMarkdown.Build()
}

// getCardTitle returns title with internal link and date to title in configured format.
func (m *Markdown) getCardTitle(card *entity.Card) string {
	title := card.Name

	switch m.config.Markdown.Format {
	case "markdown":
		title = fmt.Sprintf("[%s](cards/%s.md)", card.Name, card.Name)

		if !card.Due.IsZero() && m.config.Markdown.AddDateToCards {
			title += " `" + card.Due.Format("2006-01-02") + "`"
		}
	case "obsidian":
		title = "[[" + card.Name + "]]"

		if !card.Due.IsZero() && m.config.Markdown.AddDateToCards {
			title += " @{" + card.Due.Format("2006-01-02") + "}"
		}
	}

	return title
}

// addCardFooter adds footer with date and tags.
func (m *Markdown) addCardFooter(cardMarkdown *markdown.Markdown, card *entity.Card) {
	if m.config.Markdown.Footer != "" {
		date := ""
		if !card.Due.IsZero() {
			date = card.Due.Format("2006-01-02")
		}

		var footer string
		if strings.Contains(m.config.Markdown.Footer, "%s") {
			footer = fmt.Sprintf(m.config.Markdown.Footer, date)
		} else {
			footer = m.config.Markdown.Footer
		}

		cardMarkdown.PlainText(footer)
	}
}

func (m *Markdown) prepareDataDirectory() error {
	if m.config.Markdown.ResultDirectory == "" {
		return ErrNotExistResultFolder
	}

	m.logger.Infof("markdown: preparing data directory: %s", m.config.Markdown.ResultDirectory)

	if err := os.RemoveAll(m.config.Markdown.ResultDirectory + "/"); err != nil {
		return err
	}

	if err := files.MkdirAll(m.config.Markdown.ResultDirectory + "/cards"); err != nil {
		return err
	}

	return nil
}
