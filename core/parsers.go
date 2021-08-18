package core

import (
	"strings"

	"gopkg.in/yaml.v2"
)

type Metadata struct {
	Source string
	Url    string
}

func ParseMetadata(input string) (err error, metadata Metadata) {
	metadata = Metadata{}

	err = yaml.Unmarshal([]byte(input), &metadata)

	return
}

func ParseMarkdownBody(input string) (err error, tilbits []Tilbit) {
	texts := strings.Split(input, "\n\n")

	for _, text := range texts {
		text = strings.Trim(text, " \n")
		tilbits = append(tilbits, Tilbit{text, TilbitData{}})
	}

	return
}

func ParseMarkdownFile(fileContent string) (error, tilbits []Tilbit) {

	return
}
