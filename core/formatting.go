package core

import (
	"fmt"
	"time"

	"github.com/mitchellh/go-wordwrap"
)

func GetBitString(tilbit Tilbit) (str string, err error) {
	text := wordwrap.WrapString(tilbit.Text, 120)

	footer := fmt.Sprintf("   -- %s", tilbit.Data.Source)
	if tilbit.Data.Author != "" {
		footer += ", " + tilbit.Data.Author
	}
	if tilbit.Data.AddedOn != "" {
		footer += " (" + tilbit.Data.AddedOn + ")"
	}
	println(fmt.Sprintf("%s", tilbit.Data))

	str = text + "\n" + footer
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
