package core

import (
	"testing"

	"github.com/MarvinJWendt/testza"
)

func repo() (r *LocalSourcesRepository, tilIds []string) {
	var tilbits []Tilbit
	tilbits = append(tilbits, Tilbit{"One", SourceMetadata{}, SourceLocation{}})
	tilbits = append(tilbits, Tilbit{"Two", SourceMetadata{}, SourceLocation{}})
	tilbits = append(tilbits, Tilbit{"Three", SourceMetadata{}, SourceLocation{}})
	for _, t := range tilbits {
		tilIds = append(tilIds, t.Id())
	}
	return NewLocalSourcesRepository(tilbits), tilIds
}

func Benchmark_AllTilbits(b *testing.B) {
	for n := 0; n < b.N; n++ {
		NewLocalSourcesRepository(nil)
	}
}

func TestByQuery(t *testing.T) {
	repo, tilIds := repo()

	allBits, err := repo.ByQuery("all")

	testza.AssertNil(t, err)
	testza.AssertEqualValues(t, 3, len(allBits))
	testza.AssertContains(t, allBits, repo.loadedTilbits[0])
	testza.AssertContains(t, allBits, repo.loadedTilbits[1])
	testza.AssertContains(t, allBits, repo.loadedTilbits[2])

	randomBit, err := repo.ByQuery("random")

	testza.AssertNil(t, err)
	testza.AssertEqualValues(t, 1, len(randomBit))
	testza.AssertContains(t, tilIds, randomBit[0].Id())
	println(tilIds)

	idBit, err := repo.ByQuery(tilIds[0])

	testza.AssertNil(t, err)
	testza.AssertEqualValues(t, 1, len(idBit))
	testza.AssertEqualValues(t, idBit[0].Id(), tilIds[0])
}
