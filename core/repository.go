package core

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
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

func parseFile(fileString string) (tilbits []Tilbit) {
	fileParts := strings.Split(fileString, ".")
	extension := fileParts[len(fileParts)-1]

	data, err := ioutil.ReadFile(fileString)
	if err != nil {
		panic(err)
	}

	if extension == "md" {
		err, tilbits, _ = ParseMarkdownFile(string(data), fileString)
	} else if strings.Contains(fileString, "My Clippings.txt") {
		err, tilbits = ParseKindleClippingsFile(string(data), fileString)
	} else {
		err, tilbits = ParseTextFile(string(data), fileString)
	}
	return
}

type Repository interface {
	All() ([]Tilbit, error)
	ById(hash string) (Tilbit, error)
	ByIds(hashes []string) ([]Tilbit, error)
	ByQuery(hash string) ([]Tilbit, error)
	Create(tilbit Tilbit) (*Tilbit, error)
	Seed(tilbits []Tilbit) error
	// Update(id int64, updated Tilbit) (*Tilbit, error)
	// Delete(id int64) error
	Setup() error
}

type LocalSourcesRepository struct {
	loadedTilbits []Tilbit
}

func NewLocalSourcesRepository() Repository {
	return &LocalSourcesRepository{}
}

func (r *LocalSourcesRepository) ByIds(hashes []string) (foundBits []Tilbit, err error) {
	foundMap := map[string]Tilbit{}

	for _, hash := range hashes {
		for _, tilbit := range r.loadedTilbits {
			if strings.Contains(tilbit.Hash(), hash) {
				foundMap[hash] = tilbit
				break
			}
		}
		_, found := foundMap[hash]
		if !found {
			return []Tilbit{}, errors.New(fmt.Sprintf("Tilbit for id=[%s] not found", hash))
		}
	}
	for _, item := range foundMap {
		foundBits = append(foundBits, item)
	}

	return foundBits, nil
}

func (r *LocalSourcesRepository) ById(hash string) (tilbit Tilbit, err error) {
	tilbits, err := r.ByIds([]string{hash})
	if len(tilbits) > 0 {
		tilbit = tilbits[0]
	}

	return
}

func (r *LocalSourcesRepository) ByQuery(query string) (tilbits []Tilbit, err error) {
	if query == "all" {
		tilbits = r.loadedTilbits
	} else if query == "random" {
		randTil := getRandomBit(r.loadedTilbits)
		tilbits = append(tilbits, randTil)
	} else {
		ids := ParseIdsFromString(query)

		var err error
		tilbits, err = r.ByIds(ids)
		return tilbits, err
	}
	return
}

func (r *LocalSourcesRepository) Seed(tilbits []Tilbit) error {
	if len(tilbits) == 0 {
		tilbits = importAllSources()
	}
	r.loadedTilbits = tilbits
	return nil
}

func (r *LocalSourcesRepository) All() (tilbits []Tilbit, err error) {
	return r.loadedTilbits, nil
}
func (r *LocalSourcesRepository) Create(tilbit Tilbit) (*Tilbit, error) {
	return nil, nil
}
func (r *LocalSourcesRepository) Setup() error {
	return nil
}

func importAllSources() (tilbits []Tilbit) {
	var sources []Source

	files, err := ioutil.ReadDir(FileRepositoryDir())
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		name := FileRepositoryDir() + file.Name()

		if !file.IsDir() {
			bits := parseFile(name)
			var source Source
			source.Tilbits = bits
			source.Uri = name

			sources = append(sources, source)
		}
	}
	for _, source := range sources {
		tilbits = append(tilbits, source.Tilbits...)
	}

	return
}
