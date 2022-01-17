package core

import (
	"testing"

	"github.com/MarvinJWendt/testza"
)

func tb(name string) (tilbit Tilbit) {
	return Tilbit{name, SourceMetadata{}, SourceLocation{}, -1}
}

func seedBits() (tilbits []Tilbit) {
	tilbits = append(tilbits, tb("One"))
	tilbits = append(tilbits, tb("Two"))
	tilbits = append(tilbits, tb("Three"))
	return
}

func repo() (r Repository) {
	return NewLocalSourcesRepository(seedBits())
}

func Benchmark_AllTilbits(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewLocalSourcesRepository(nil)
	}
}

func TestByQuery(t *testing.T) {
	repo := repo()

	repoBits := seedBits()

	allBits, err := repo.ByQuery("all")

	testza.AssertNil(t, err)
	testza.AssertEqualValues(t, 3, len(allBits))
	testza.AssertContains(t, allBits, repoBits[0])
	testza.AssertContains(t, allBits, repoBits[1])
	testza.AssertContains(t, allBits, repoBits[2])

	randomBit, err := repo.ByQuery("random")

	testza.AssertNil(t, err)
	testza.AssertEqualValues(t, 1, len(randomBit))
	testza.AssertContains(t, repoBits, randomBit[0])

	idBit, err := repo.ByQuery(repoBits[0].Hash())

	testza.AssertNil(t, err)
	testza.AssertEqualValues(t, 1, len(idBit))
	testza.AssertEqualValues(t, idBit[0].Hash(), repoBits[0].Hash())
}
