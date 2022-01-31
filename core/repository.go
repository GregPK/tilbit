package core

import (
	"math/rand"
	"os"
	"time"
)

type Source struct {
	Uri     string
	Tilbits []Tilbit
}

func FileRepositoryDir() string {
	homeDir, _ := os.UserHomeDir()
	return homeDir + "/.config/tilbit/data/"
}

func getRandomBit(tilbits []Tilbit) (randomTilbit Tilbit) {
	rand.Seed(time.Now().UnixNano())
	randomTilbit = tilbits[rand.Intn(len(tilbits))]
	return
}

type Repository interface {
	All() ([]Tilbit, error)
	ById(hash string) (Tilbit, error)
	ByIds(hashes []string) ([]Tilbit, error)
	ByQuery(hash string) ([]Tilbit, error)
	Create(tilbit Tilbit) (*Tilbit, error)
	Import(tilbits []Tilbit) error
	// Update(id int64, updated Tilbit) (*Tilbit, error)
	// Delete(id int64) error
	Setup() error
	SetSourceURI(source string) error
}

func GetTilbits(sources []string, query string, inputTilbits []Tilbit) (tilbits []Tilbit, err error) {
	repo := NewSQLiteRepository()
	if len(sources) == 0 && len(inputTilbits) > 0 { // mostly for input tilbit testing
		repo.Import(inputTilbits)
		sources = append(sources, "")
	} else {
		if len(sources) == 0 {
			sources = append(sources, FileRepositoryDir())
		}
		for _, source := range sources {
			bits, _ := ParseSource(source)
			repo.Import(bits)
		}
	}

	repoBits, repoErr := repo.ByQuery(query)
	return repoBits, repoErr
}
