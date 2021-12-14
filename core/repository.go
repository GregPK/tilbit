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

func AllTilbits() (tilbits []Tilbit) {
	sources := LoadSources()

	for _, source := range sources {
		tilbits = append(tilbits, source.Tilbits...)
	}
	return
}

func ByQuery(query string, inputTilbits []Tilbit) (tilbits []Tilbit, err error) {
	if len(inputTilbits) == 0 {
		inputTilbits = AllTilbits()
	}

	if query == "all" {
		tilbits = inputTilbits
	} else if query == "random" {
		randTil := getRandomBit(inputTilbits)
		tilbits = append(tilbits, randTil)
	} else {
		ids := ParseIdsFromString(query)

		var err error
		tilbits, err = ByIds(ids, inputTilbits)
		return tilbits, err
	}
	return
}

func ById(hash string, inputTilbits []Tilbit) (tilbit Tilbit, err error) {
	tilbits, err := ByIds([]string{hash}, inputTilbits)
	if len(tilbits) > 0 {
		tilbit = tilbits[0]
	}

	return
}

func ByIds(hashes []string, inputTilbits []Tilbit) (foundBits []Tilbit, err error) {
	foundMap := map[string]Tilbit{}

	for _, hash := range hashes {
		for _, tilbit := range inputTilbits {
			if strings.Contains(tilbit.Id(), hash) {
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

func LoadSources() (sources []Source) {
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
	return
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
