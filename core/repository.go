package core

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
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

func ById(hash string) (tilbit Tilbit, err error) {
	tilbits, err := ByIds([]string{hash})
	if len(tilbits) > 0 {
		tilbit = tilbits[0]
	}

	return
}

func ByIds(hashes []string) (foundBits []Tilbit, err error) {
	tilbits := AllTilbits()

	foundMap := map[string]Tilbit{}

	for _, hash := range hashes {
		for _, tilbit := range tilbits {
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
