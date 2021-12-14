package core

import (
	"testing"

	"github.com/MarvinJWendt/testza"
)

func tilbits() (tilbits []Tilbit, tilIds []string) {
	tilbits = append(tilbits, Tilbit{"One", SourceMetadata{}, SourceLocation{}})
	tilbits = append(tilbits, Tilbit{"Two", SourceMetadata{}, SourceLocation{}})
	tilbits = append(tilbits, Tilbit{"Three", SourceMetadata{}, SourceLocation{}})
	for _, t := range tilbits {
		tilIds = append(tilIds, t.Id())
	}
	return
}

func Benchmark_AllTilbits(b *testing.B) {
	for n := 0; n < b.N; n++ {
		AllTilbits()
	}
}

func TestByQuery(t *testing.T) {
	tilbits, tilIds := tilbits()

	allBits, err := ByQuery("all", tilbits)

	testza.AssertNil(t, err)
	testza.AssertEqualValues(t, 3, len(allBits))
	testza.AssertContains(t, allBits, tilbits[0])
	testza.AssertContains(t, allBits, tilbits[1])
	testza.AssertContains(t, allBits, tilbits[2])

	randomBit, err := ByQuery("random", tilbits)

	testza.AssertNil(t, err)
	testza.AssertEqualValues(t, 1, len(randomBit))
	testza.AssertContains(t, tilIds, randomBit[0].Id())
	println(tilIds)

	idBit, err := ByQuery(tilIds[0], tilbits)

	testza.AssertNil(t, err)
	testza.AssertEqualValues(t, 1, len(idBit))
	testza.AssertEqualValues(t, idBit[0].Id(), tilIds[0])
}
