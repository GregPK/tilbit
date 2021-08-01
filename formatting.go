
package main

import (
	"time"
	"fmt"

	"github.com/mitchellh/go-wordwrap"
)

func getBitString(tilbit Tilbit) (str string, err error) {
	text := wordwrap.WrapString(tilbit.Text, 120)
	footer := fmt.Sprintf("   -- %s (%s)", tilbit.Data.Source, tilbit.Data.AddedOn)
	str = fmt.Sprintf("%s\n%s", text, footer)
	return
}
func makeTilLine(content string, source string) (tilLine string) {
	ISO8601 := "2006-_2-_1"
	addedOn := time.Now().Format(ISO8601)

	tilLine = fmt.Sprintf("%s, {\"source\": \"%s\", addedOn:\"%s\"}\n\n", content, source, addedOn)
	return
}