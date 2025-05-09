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

const cardFooterTemplate = `---

- date: %s
- tags: #todo`

type Markdown struct {
	config *config.Config
	logger *logger.Logger
}

func New(config *config.Config, logger *logger.Logger) *Markdown {
	m := Markdown{
		config: config,
		logger: logger,
	}

	return &m
}

func (m *Markdown) CreateMarkdown(boardName string) error {
	file, err := os.ReadFile(boardName)
	if err != nil {
		return err
	}

	var dash entity.Dashboard
	if err = json.Unmarshal(file, &dash); err != nil {
		return err
	}

	// Prepare indexes

	sort.Slice(dash.Cards, func(i, j int) bool {
		return dash.Cards[i].Due.After(dash.Cards[j].Due)
	})

	cards := make(map[string]entity.Card)
	for _, card := range dash.Cards {
		if !card.Closed {
			cards[card.ID] = card
		}
	}

	checklists := make(map[string]entity.Checklist)
	for _, check := range dash.Checklists {
		checklists[check.ID] = check
	}

	// Create main markdown

	f, err := os.Create(m.config.App.HomeDirectory + "/.tmp/data/response/" + dash.Name + ".md")
	if err != nil {
		return err
	}
	defer f.Close()

	mdBoard := markdown.NewMarkdown(f)

	for _, list := range dash.Lists {
		if list.Closed {
			continue
		}

		mdBoard.H2(list.Name)

		for _, card := range dash.Cards {
			if card.IDList == list.ID {
				cardName := "[[" + card.Name + "]]"

				if !card.Due.IsZero() {
					cardName += " @{" + card.Due.Format("2006-01-02") + "}"
				}

				mdBoard.CheckBox([]markdown.CheckBoxSet{
					{Checked: card.DueComplete, Text: cardName},
				})
			}
		}

		mdBoard.PlainText("")
	}

	if err := mdBoard.Build(); err != nil {
		return err
	}

	// Create checklists markdowns

	for _, checklist := range dash.Checklists {
		card, ok := cards[checklist.IDCard]
		if !ok || card.Closed {
			continue
		}

		f, err := os.Create(m.config.App.HomeDirectory + "/.tmp/data/response/tododata/" + card.Name + ".md")
		if err != nil {
			return err
		}

		cardBoard := markdown.NewMarkdown(f)

		cardBoard.H1(card.Name)

		for _, item := range checklist.CheckItems {
			cardBoard.CheckBox([]markdown.CheckBoxSet{
				{Checked: item.State == "complete", Text: item.Name},
			})
		}

		date := ""
		if !card.Due.IsZero() {
			date = card.Due.Format("2006-01-02")
		}

		cardBoard.PlainText("").PlainText(fmt.Sprintf(cardFooterTemplate, date))

		if err := cardBoard.Build(); err != nil {
			_ = f.Close()

			return err
		}

		if err := f.Close(); err != nil {
			return err
		}
	}

	return nil
}
