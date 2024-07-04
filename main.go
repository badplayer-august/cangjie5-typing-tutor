package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

  "github.com/badplayer-august/cangjie5-typing-tutor/src/cangjie"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

var data cangjie.Dataset
var docStyle = lipgloss.NewStyle().Padding(1, 2, 1, 2)

func main() {
  data = cangjie.LoadCangjie()
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

type Model struct {
	Selected   cangjie.CharInfo
	UserInput  [5]rune
	ShowAsnwer bool
}

func RandomSelectQuestion() cangjie.CharInfo {
	r := rand.Intn(100)
	switch {
	case r < 10:
		return data.C[rand.Intn(len(data.C))]
	case r < 30:
		return data.B[rand.Intn(len(data.B))]
	default:
		return data.A[rand.Intn(len(data.A))]
	}
}

func initialModel() Model {
	return Model{
		Selected:   RandomSelectQuestion(),
		UserInput:  [5]rune{},
		ShowAsnwer: false,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			m.Selected = RandomSelectQuestion()
			m.UserInput = [5]rune{}
		}
	}
	return m, nil
}

func (m Model) View() string {
	physicalWidth, _, _ := term.GetSize(int(os.Stdout.Fd()))
	doc := strings.Builder{}

	// Title
	{
		row := lipgloss.NewStyle().Align(lipgloss.Center).Render("Cangjie Type Tutor!!")
		doc.WriteString(row + "\n\n")
	}
	if physicalWidth > 0 {
		docStyle = docStyle.MaxWidth(physicalWidth)
	}

  s := docStyle.Render(doc.String())
	s += "Cangjie Type Tutor\n"
	s += fmt.Sprintf("Char: %s\n", m.Selected.Char)
	s += fmt.Sprintf("UserInput: %v\n", m.UserInput)
	if m.ShowAsnwer {
		s += fmt.Sprintf("Key: %v\n", m.Selected.Key)
		s += fmt.Sprintf("Full Key: %v\n", StrFormat(m.Selected.Key[0], false))
		s += fmt.Sprintf("Name: %v\n", StrFormat(m.Selected.Key[0], true))
	}
	s += "\nPress ctrl-c to quit.\n"
	return s
}

func StrFormat(s string, to_name bool) string {
	runes := []rune(s)
	ss := make([]string, len(runes))
	if to_name {
		for i, r := range runes {
			ss[i] = cangjie.KeyToName[r]
		}
	} else {
		for i, r := range runes {
			ss[i] = string(r + 65216)
		}
	}
	return strings.Join(ss, " ")
}
