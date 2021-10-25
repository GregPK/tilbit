package core

import (
	"testing"

	"github.com/MarvinJWendt/testza"
)

func textFile() string {
	return `
First tilbit {"addedOn": "2021-04-13", "source": "https://example.com"}

Second tilbit. {"addedOn": "2021-03-29", "url":"https://gregpk.com"}
`
}

func basicMarkdown() string {
	return `
**First TIL:** Content 1
* First item
* Second item

**Second TIL:** Content 2
`
}

func frontmatter() string {
	return `
---
source: Source title
url: https://gregpk.com
---
`
}

func TestMetadataParser(t *testing.T) {
	err, m := ParseMetadata(frontmatter())
	testza.AssertEqual(t, m.Source, "Source title")
	testza.AssertEqual(t, m.Url, "https://gregpk.com")
	testza.AssertNil(t, err)
}

func TestMarkdownParserParagraphs(t *testing.T) {
	test := `
**First TIL:** Content 1

**Second TIL:** Content 2
`

	err, tilbits := ParseMarkdownBody(test)
	testza.AssertEqual(t, len(tilbits), 2)
	testza.AssertEqual(t, tilbits[0].Text, `**First TIL:** Content 1`)
	testza.AssertEqual(t, tilbits[1].Text, `**Second TIL:** Content 2`)
	testza.AssertNil(t, err)
}

func TestMarkdownParserBody(t *testing.T) {
	err, tilbits := ParseMarkdownBody(basicMarkdown())
	testza.AssertEqual(t, len(tilbits), 2)
	testza.AssertEqual(t, tilbits[0].Text, "**First TIL:** Content 1\n* First item\n* Second item")
	testza.AssertEqual(t, tilbits[1].Text, `**Second TIL:** Content 2`)
	testza.AssertNil(t, err)
}

// ---
func TestMarkdownFileBasic(t *testing.T) {
	fileContent := frontmatter() + "\n" + basicMarkdown()

	err, tilbits, metadata := ParseMarkdownFile(fileContent)

	testza.AssertEqual(t, metadata.Source, "Source title")
	testza.AssertEqual(t, metadata.Url, "https://gregpk.com")

	testza.AssertEqual(t, len(tilbits), 2)
	testza.AssertEqual(t, tilbits[0].Text, "**First TIL:** Content 1\n* First item\n* Second item")
	testza.AssertEqual(t, tilbits[1].Text, `**Second TIL:** Content 2`)
	testza.AssertEqual(t, tilbits[0].Data.Source, "Source title")
	testza.AssertEqual(t, tilbits[1].Data.Source, "Source title")
	testza.AssertEqual(t, tilbits[0].Data.Url, "https://gregpk.com")

	testza.AssertNil(t, err)
}

func TestTextFile(t *testing.T) {
	fileContent := textFile()

	err, tilbits := ParseTextFile(fileContent)

	testza.AssertEqual(t, err, nil)

	testza.AssertEqual(t, len(tilbits), 2)

	testza.AssertEqual(t, tilbits[0].Text, "First tilbit")
	testza.AssertEqual(t, tilbits[0].Data.Source, "https://example.com")
	testza.AssertEqual(t, tilbits[0].Data.AddedOn, "2021-04-13")

	testza.AssertEqual(t, tilbits[1].Text, "Second tilbit.")
	testza.AssertEqual(t, tilbits[1].Data.Source, "")
	testza.AssertEqual(t, tilbits[1].Data.AddedOn, "2021-03-29")
	testza.AssertEqual(t, tilbits[1].Data.Url, "https://gregpk.com")
}

func kindleClippingsFile() string {
	return `
Source Title (Author1, Name)
- Your Highlight on page 110 | location 1681-1682 | Added on Sunday, 15 August 2021 14:21:39

Short highlight.
==========
Source Title
- Your Highlight on page 189 | location 2892-2893 | Added on Sunday, 15 August 2021 18:27:59

Longer Tilbit
Perhaps longer than one line.
==========
`
}
func TestKindleClippingsFile(t *testing.T) {
	fileContent := kindleClippingsFile()

	err, tilbits := ParseKindleClippingsFile(fileContent)

	testza.AssertEqual(t, err, nil)

	testza.AssertEqual(t, len(tilbits), 2)

	testza.AssertEqual(t, tilbits[0].Text, "Short highlight.")
	testza.AssertEqual(t, tilbits[0].Data.Source, "Source Title (Author1, Name)")
	// testza.AssertEqual(t, tilbits[0].Data.AddedOn, "2021-04-13")

	testza.AssertEqual(t, tilbits[1].Text, "Longer Tilbit\nPerhaps longer than one line.")
	testza.AssertEqual(t, tilbits[1].Data.Source, "Source Title")
	// testza.AssertEqual(t, tilbits[1].Data.AddedOn, "2021-03-29")
}
