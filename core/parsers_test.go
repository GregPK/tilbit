package core

import (
	"testing"

	"github.com/MarvinJWendt/testza"
)

func TestMetadataParser(t *testing.T) {
	test := `
source: 40 powerful concepts
url: https://twitter.com/g_s_bhogal/status/1225561131122597896
`

	err, m := ParseMetadata(test)
	testza.AssertEqual(t, m.Source, "40 powerful concepts")
	testza.AssertEqual(t, m.Url, "https://twitter.com/g_s_bhogal/status/1225561131122597896")
	testza.AssertNil(t, err)
}

func TestMarkdownParserParagraphs(t *testing.T) {
test := `
**First TIL:** Content 1

**Second TIL:** Content 2
`
	err, tilbits := ParseMarkdown(test)
	testza.AssertEqual(t, len(tilbits), 2)
	testza.AssertNil(t, err)
}