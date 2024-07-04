package ui

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

var (
	// General
	docStyle   = lipgloss.NewStyle().Padding(1, 0, 1, 0)
	titleStyle = lipgloss.NewStyle().Align(lipgloss.Center)

	// Answer
	defaultAnswerStyle = lipgloss.AdaptiveColor{Light: "#4c4f69", Dark: "#cad3f5"}
	correctAnswerStyle = lipgloss.AdaptiveColor{Light: "#40a02b", Dark: "#a6e3a1"}
	wrongAnswerStyle   = lipgloss.AdaptiveColor{Light: "#d20f39", Dark: "#f38ba8"}
)

func Render(title string, char string) string {
	physicalWidth, _, _ := term.GetSize(int(os.Stdout.Fd()))
	doc := strings.Builder{}

	if physicalWidth > 0 {
		docStyle = docStyle.MaxWidth(physicalWidth)
	}

  fmt.Println(physicalWidth)
  // title
  {
    row := titleStyle.Width(physicalWidth).Render(title)
    doc.WriteString(row + "\n\n")
  }

  // char
  {
    row := titleStyle.Width(physicalWidth).Height(10).Foreground(correctAnswerStyle).Render(char)
    doc.WriteString(row + "\n\n")
  }

	return docStyle.Render(doc.String())
}
