package core

import (
	"crypto/md5"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/mitchellh/go-wordwrap"
	"golang.org/x/term"
)

func GetBitString(tilbit Tilbit, box bool) (str string, err error) {
	box = true

	if box {
		printBox(tilbit)
	} else {
		printString(tilbit)
	}

	return
}

func printBox(tilbit Tilbit) {
	text, wrapWidth := wrapText(tilbit.Text)
	footerId := printFooter(tilbit, true, false)
	footer := printFooter(tilbit, false, false)
	const boxMargin = 10

	Box := box.New(box.Config{Px: 1, Py: 0, Type: "Single", Color: randomColor(), TitlePos: "Top"})

	if len(footerId) < int(wrapWidth)-boxMargin {
		Box.Print(footerId, text)
	} else if len(footer) < int(wrapWidth)-boxMargin {
		Box.Print(footer, text)
	} else {
		wrappedFooter, _ := wrapText(footer)
		Box.Print(tilbit.Id(), text+"\n"+wrappedFooter)
	}

	return
}

func printString(tilbit Tilbit) {
	text := tilbit.Text + "\n" + printFooter(tilbit, true, true)
	text, _ = wrapText(text)

	fmt.Print(text)

	return
}

// Wraps and pads the body text based on the terminal size
// Returns the result text and the wrapped size
func wrapText(text string) (wrapped string, wrapWidth uint) {
	termSize, _, _ := term.GetSize(int(os.Stdin.Fd()))
	wrapWidth = uint(math.Min(float64(120), float64(termSize)-10))

	if len(text) > int(wrapWidth) {
		wrapped = wordwrap.WrapString(text, wrapWidth)
	} else {
		paddingLen := int(wrapWidth) - len(text) - 1
		wrapped = text + strings.Repeat(" ", paddingLen)
	}

	return
}

func printFooter(tilbit Tilbit, appendId bool, indent bool) (footer string) {
	if indent {
		footer += "   -- "
	}

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

func randomColor() string {
	rand.Seed(time.Now().UnixNano())
	colors := []string{"Black", "Blue", "Red", "Green", "Yellow", "Cyan", "Magenta", "White"}
	randomIndex := rand.Intn(len(colors))
	color := colors[randomIndex]
	if rand.Intn(2) == 1 || true {
		color = "Hi" + color
	}
	return color
}

func MakeTilLine(content string, source string) (tilLine string) {
	addedOn := isoDate(time.Now())

	tilLine = fmt.Sprintf("%s, {\"source\": \"%s\", addedOn:\"%s\"}\n\n", content, source, addedOn)
	return
}

func (t Tilbit) Id() string {
	return t.Hash()[:8]
}

func (t Tilbit) Hash() string {
	h := md5.New()
	io.WriteString(h, t.Text)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func isoDate(timeToFormat time.Time) string {
	ISO8601 := "2006-02-01"
	return timeToFormat.Format(ISO8601)
}
