package main

import (
	"bytes"
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"os"
)

var history [30]string
var historyPointer int = 0
var historyCount int = 0

type (
	errMsg error
)

func main() {
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func handle(message string) string {
	message = string(bytes.Trim([]byte(message), "\x00"))
	result, errors := send(message)

	if historyCount == len(history) {
		history = [30]string(history[1:30])
	} else {
		history[historyCount] = message
		historyCount++
	}

	if errors != nil {
		return errors.Error()
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
			tea.ClearScreen()
			msg := m.textInput.Value()
			m.textInput.SetValue("")
			historyPointer = historyCount
			ans := handle(msg)
			return m, tea.Printf(ans)
		case tea.KeyEscape, tea.KeyCtrlC:
			os.Exit(0)
			return m, tea.Quit
		case tea.KeyUp:
			if historyPointer <= 0 {
				historyPointer = historyCount
			}
			his := history[historyPointer]

			m.textInput.SetValue(his)

			historyPointer--
			return m, nil
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
