package core

import (
	"testing"

	"github.com/MarvinJWendt/testza"
)

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
source: Source title
url: https://gregpk.com
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