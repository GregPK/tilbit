package core

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v2"
)

const (
	frontmatterBreak = "---"
)

func ParseMetadata(input string) (err error, metadata SourceMetadata) {
	metadata = SourceMetadata{}

	input = strings.ReplaceAll(input, frontmatterBreak, "")
	input = strings.Trim(input, "\n ")
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

func ParseMarkdownFile(fileContent string) (err error, tilbits []Tilbit, metadata SourceMetadata) {
	frontmatter := ""
	body := ""

	parts := strings.Split(fileContent, frontmatterBreak)
	if len(parts) > 1 {
		frontmatter = parts[1]
		body = parts[2]
	} else {
		body = parts[0]
	}

	err, metadata = ParseMetadata(strings.TrimSpace(frontmatter))
	err, tilbits = ParseMarkdownBody(strings.TrimSpace(body))

	fmt.Printf("FM: [%v]\nBD: [%v]", metadata, tilbits)

	if metadata.Source != "" || metadata.Url != "" {
		for i, tilbit := range tilbits {
			if metadata.Source != "" {
				tilbit.Data.Source = metadata.Source
			}
			if metadata.Url != "" {
				tilbit.Data.Url = metadata.Url
			}
			tilbits[i] = tilbit
		}
	}

	return
}
