package core

import (
	"fmt"
	"math"
	"time"

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
func wrapText(text string) (wrapped string, wrapWidth uint) {
	termSize := pterm.GetTerminalWidth()
	if termSize < 1 { // probably running headless
		termSize = defTermSize
	}
	wrapWidth = uint(math.Min(float64(defTermSize), float64(termSize)-3))

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
