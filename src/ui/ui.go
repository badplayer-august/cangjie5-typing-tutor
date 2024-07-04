package ui

import (
	"os"
	"strings"

	"github.com/badplayer-august/cangjie5-typing-tutor/src/cangjie"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

var (
	// General
	docStyle   = lipgloss.NewStyle().Align(lipgloss.Center).Padding(1, 2, 1, 2)
	titleStyle = lipgloss.NewStyle()
	helpStyle  = lipgloss.NewStyle().Italic(true).Faint(true)

	// Answer
	correctAnswerColor = lipgloss.Color("2")
	wrongAnswerColor   = lipgloss.Color("1")
)

func Render(title string, char string, input string) string {
	physicalWidth, _, _ := term.GetSize(int(os.Stdout.Fd()))
	docStyle = docStyle.Width(physicalWidth)
	doc := strings.Builder{}

	if physicalWidth > 0 {
		docStyle = docStyle.MaxWidth(physicalWidth)
	}

	// title
	{
		row := titleStyle.Render(title)
		doc.WriteString(row + "\n\n")
	}

	// char
	{
		row := titleStyle.Foreground(wrongAnswerColor).Render(char)
		doc.WriteString(row + "\n\n")
	}

	// kays and names
	{
		keys, names := formatDecomposition(input)
    row := lipgloss.JoinVertical(lipgloss.Center, names, keys)
		doc.WriteString(row + "\n\n")
	}

	{
		row := helpStyle.Render("Press ctrl-c to quit.")
		doc.WriteString(row + "\n\n")
	}

	return docStyle.Render(doc.String())
}

func formatDecomposition(decomp string) (string, string) {
			var (
				runes = []rune(decomp)
				keys  = make([]string, len(runes))
				names = make([]string, len(runes))
			)

			for i, r := range runes {
				keys[i] = string(r + 65216)
				names[i] = cangjie.KeyToName[r]
			}

			return strings.Join(keys, " "), strings.Join(names, " ")
		}
