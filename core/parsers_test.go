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
