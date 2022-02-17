package core

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/pterm/pterm"
)

type Style struct {
	bg string
	fg string
}

func GetRandomStyle() (style Style) {
	niceAnsiColorsLight := [10]int{209, 204, 198, 162, 42, 39, 33, 32, 31, 27}
	niceAnsiColorsDark := [21]int{226, 229, 213, 211, 209, 204, 198, 197, 196, 191, 190, 184, 162, 156, 155, 154, 118, 115, 86, 82, 47}

	styles := []Style{}
	for _, col := range niceAnsiColorsLight {
		styles = append(styles, Style{strconv.Itoa(col), "#eeeeee"})
	}
	for _, col := range niceAnsiColorsDark {
		styles = append(styles, Style{strconv.Itoa(col), "#111111"})
	}

	rand.Seed(time.Now().UnixNano())
	style = styles[rand.Intn(len(styles))]

	return
}

func Block(tilbit Tilbit) string {
	style := GetRandomStyle()

	// blockStyle := lipgloss.NewStyle().
	// 	Margin(1, 2).
	// Width(80)

	maxWidth := pterm.GetTerminalWidth() - 2

	bodyStyle := lipgloss.NewStyle().
		Align(lipgloss.Left).
		Padding(1, 2).
		Margin(0, 1).
		Background(lipgloss.Color(style.bg)).
		Foreground(lipgloss.Color(style.fg)).
		Width(maxWidth)

	var source string
	// var addedOn string
	var id string

	barStyle := lipgloss.NewStyle().
		Background(lipgloss.Color(style.bg)).
		Foreground(lipgloss.Color(style.fg)).
		Margin(0, 1).
		Padding(0, 2).
		MaxWidth(maxWidth - 10)

	titleStr := tilbit.Data.Source
	if tilbit.Data.Author != "" {
		titleStr += " – " + tilbit.Data.Author
	}

	w := lipgloss.Width
	if titleStr != "" {
		source = barStyle.Render(tilbit.Data.Source)
	}
	// if tilbit.Data.AddedOn != "" {
	// addedOn = lipgloss.NewStyle().Render(tilbit.Data.AddedOn)
	// }
	id = lipgloss.NewStyle().Render(tilbit.Id())

	styleLine := "" // style.bg + " / " + style.fg + "\n"
	body := bodyStyle.Render(styleLine + tilbit.Text)
	gap := lipgloss.NewStyle().Width(w(body) - w(source) - w(id) - 1).Render("")
	bar := lipgloss.JoinHorizontal(lipgloss.Top,
		source,
		gap,
		id,
	)

	return bar + "\n" + body + "\n\n"
}
