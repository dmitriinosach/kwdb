package api

import (
	"context"
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"kwdb/app"
	"kwdb/app/logger"
	"log"
	"os"
)

type (
	errMsg error
)

func HandleCLI() {
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func handle(message string) string {
	ctx := context.Background()
	result, err := app.HandleMessage(ctx, message)

	if err != nil {
		logger.Write(err.Error())
		return err.Error()
	}

	return result
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = ""
	ti.Focus()
	ti.CharLimit = 2048
	ti.Width = 200

	return model{
		textInput: ti,
		err:       nil,
	}
}

type model struct {
	textInput textinput.Model
	err       error
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			result := handle(m.textInput.Value())
			m.textInput.SetValue("")
			return m, tea.Println(result)
		case tea.KeyEscape:
			os.Exit(0)
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)

	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"kvdb%s\n%s",
		m.textInput.View(),
		"(esc to quit)",
	) + "\n"
}
