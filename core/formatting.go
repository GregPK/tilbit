package core

import (
	"crypto/md5"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/mitchellh/go-wordwrap"
	"golang.org/x/term"
)

func GetBitString(tilbit Tilbit, box bool) (str string, err error) {
	box = true

	if box {
		str = printBox(tilbit)
	} else {
		text, footer, _ := printString(tilbit)
		str = text + "\n" + footer
	}

	return
}

func printBox(tilbit Tilbit) (str string) {
	str = ""
	Box := box.New(box.Config{Px: 1, Py: 0, Type: "Single", Color: randomColor(), TitlePos: "Top"})
	text, footer, wrapWidth := printString(tilbit)

	if len(footer) < int(wrapWidth) {
		Box.Print(footer, text)
	} else {
		text := text + "\n" + "  " + wordwrap.WrapString(footer, wrapWidth-5)
		Box.Print("", text)
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

func printString(tilbit Tilbit) (text string, footer string, wrapWidth uint) {
	termSize, _, _ := term.GetSize(int(os.Stdin.Fd()))
	wrapWidth = uint(math.Min(float64(120), float64(termSize)-10))
	text = wordwrap.WrapString(tilbit.Text, wrapWidth)

	footer = fmt.Sprintf("   -- %s", tilbit.Data.Source)
	if tilbit.Data.Author != "" {
		footer += ", " + tilbit.Data.Author
	}
	if tilbit.Data.AddedOn != "" {
		footer += " (" + tilbit.Data.AddedOn + ")"
	}
	footer += " (id: " + tilbit.Id()[:8] + ")"

	return
}

func MakeTilLine(content string, source string) (tilLine string) {
	addedOn := isoDate(time.Now())

	tilLine = fmt.Sprintf("%s, {\"source\": \"%s\", addedOn:\"%s\"}\n\n", content, source, addedOn)
	return
}

func (t Tilbit) Id() string {
	h := md5.New()
	io.WriteString(h, t.Text)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func isoDate(timeToFormat time.Time) string {
	ISO8601 := "2006-02-01"
	return timeToFormat.Format(ISO8601)
}
