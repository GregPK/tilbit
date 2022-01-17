package core

import (
	"testing"

	"github.com/MarvinJWendt/testza"
)

const useNewDbRepo = true

func tb(name string) (tilbit Tilbit) {
	return Tilbit{name, SourceMetadata{}, SourceLocation{}, -1, ""}
}

func seedBits() (tilbits []Tilbit) {
	tilbits = append(tilbits, tb("One"))
	tilbits = append(tilbits, tb("Two"))
	tilbits = append(tilbits, tb("Three"))
	return
}

func repo() (r Repository) {
	if useNewDbRepo {
		r = NewSQLiteRepository()
	} else {
		r = NewLocalSourcesRepository()
	}
	err := r.Seed(seedBits())
	if err != nil {
		panic(err)
	}
	return
}

func hashes(tilbits []Tilbit) (hashes []string) {
	for _, tilbit := range tilbits {
		hashes = append(hashes, tilbit.Hash())
	}
	return
}

func Benchmark_AllTilbits(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewLocalSourcesRepository()
	}
}

func TestByQuery(t *testing.T) {
	repo := repo()

	repoBits := seedBits()

	allBits, err := repo.ByQuery("all")
	allHashes := hashes(allBits)

	testza.AssertNil(t, err)
	testza.AssertEqualValues(t, 3, len(allBits))
	testza.AssertContains(t, allHashes, repoBits[0].Hash())
	testza.AssertContains(t, allHashes, repoBits[1].Hash())
	testza.AssertContains(t, allHashes, repoBits[2].Hash())

	randomBit, err := repo.ByQuery("random")

	testza.AssertNil(t, err)
	testza.AssertEqualValues(t, 1, len(randomBit))
	testza.AssertContains(t, hashes(repoBits), randomBit[0].Hash())

	idBit, err := repo.ByQuery(repoBits[0].Hash())

	testza.AssertNil(t, err)
	testza.AssertEqualValues(t, 1, len(idBit))
	testza.AssertEqualValues(t, idBit[0].Hash(), repoBits[0].Hash())
}
