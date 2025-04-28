package main

import (
	"flag"
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"net"
	"os"
	"sync"
)

type (
	errMsg error
)

// TODO: как это работает
//
//go:generate stringer -type=cliConfig
var cliConfig struct {
	connectionHost string
	connectionPort string
}

var cliHistory = new(history)

func main() {
	// локальный IP
	//locIp := helper.LocalIp()

	// Регистрируем флаги и связываем их с полями структуры config
	flag.StringVar(&cliConfig.connectionHost, "host", "192.168.1.96", "хост для подключения")
	flag.StringVar(&cliConfig.connectionPort, "port", "712", "порт для подключения")

	// Парсим аргументы командной строки
	flag.Parse()

	status := handle("status")

	fmt.Println("подключено к " + net.JoinHostPort(cliConfig.connectionHost, cliConfig.connectionPort))
	fmt.Println(status)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()

		p := tea.NewProgram(initialModel())
		if _, err := p.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	wg.Wait()
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
			inpt := m.textInput.Value()
			m.textInput.SetValue("")
			if inpt != "" {
				cliHistory.Push(inpt)
				ans := handle(inpt)
				return m, tea.Println(ans)
			}
		case tea.KeyEscape, tea.KeyCtrlC:
			os.Exit(0)
			return m, tea.Quit
		case tea.KeyUp:

			hisCmd := cliHistory.Prev()

			m.textInput.SetValue(hisCmd)

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
