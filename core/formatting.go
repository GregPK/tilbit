package core

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"github.com/eidolon/wordwrap"
	"github.com/pterm/pterm"
	"gopkg.in/yaml.v2"
)

func GetBitString(tilbit Tilbit, format string) (str string, err error) {
	if format == "box" {
		printBox(tilbit)
	} else if format == "yaml" {
		str = printYaml(tilbit)
	} else {
		str = printString(tilbit)
	}

	return
}

func printBox(tilbit Tilbit) {
	rand.Seed(time.Now().UnixNano())

	color := strconv.Itoa(rand.Intn(15))
	dialogBoxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(color)).
		Padding(0, 1).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).
		BorderBottom(true)

	_, width := wrapText(tilbit.Text)
	footerId := printFooter(tilbit, true, false)
	// footer := printFooter(tilbit, false, false)

	style := `
	{
		"document": {
			"block_prefix": "",
			"block_suffix": "",
			"margin": 0
		}
	}
	`

	r, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(0),
		glamour.WithStylesFromJSONBytes([]byte(style)),
	)
	out, _ := r.Render(tilbit.Text)

	question := lipgloss.NewStyle().Width(width - 1).Align(lipgloss.Left).Render(out + "\n" + footerId)

	dialog := lipgloss.Place(width, -1,
		lipgloss.Left, lipgloss.Left,
		dialogBoxStyle.Render(question),
	)

	fmt.Print(dialog)
}

func printPtermBox(tilbit Tilbit) {
	text, _ := wrapText(tilbit.Text)
	footerId := printFooter(tilbit, true, false)
	// footer := printFooter(tilbit, false, false)

	pterm.DefaultBox.WithTitle(footerId).WithTitleBottomRight().WithRightPadding(0).WithBottomPadding(0).Println(text)
}

func printString(tilbit Tilbit) (text string) {
	text = tilbit.Text + "\n" + printFooter(tilbit, true, true)

	return
}

func printYaml(tilbit Tilbit) (text string) {
	data, e := yaml.Marshal(&tilbit)
	if e != nil {
		text = ""
	} else {
		text = string(data)
	}
	return
}

var defTermSize = 120

// Wraps and pads the body text based on the terminal size
// Returns the result text and the wrapped size
func wrapText(text string) (wrapped string, wrapWidth int) {
	termSize := pterm.GetTerminalWidth()
	if termSize < 1 { // probably running headless
		termSize = defTermSize
	}
	wrapWidth = int(math.Min(float64(defTermSize), float64(termSize)-3))

	if len(text) >= int(wrapWidth) {
		wrapper := wordwrap.Wrapper(int(wrapWidth), true)
		wrapped = wrapper(text)
	} else {
		wrapped = text
	}

	return
}

func printFooter(tilbit Tilbit, appendId bool, indent bool) (footer string) {
	if tilbit.Data.Author != "" {
		footer += tilbit.Data.Author
	}
	if tilbit.Data.Source != "" {
		footer += " - " + tilbit.Data.Source
	}
	if tilbit.Data.AddedOn != "" {
		footer += " (" + tilbit.Data.AddedOn + ")"
	}
	if appendId {
		footer += " (id: " + tilbit.Id() + ")"
	}

	return
}

func MakeTilLine(content string, source string) (tilLine string) {
	addedOn := isoDate(time.Now())

	tilLine = fmt.Sprintf("%s, {\"source\": \"%s\", addedOn:\"%s\"}\n\n", content, source, addedOn)
	return
}

func isoDate(timeToFormat time.Time) string {
	ISO8601 := "2006-02-01"
	return timeToFormat.Format(ISO8601)
}
