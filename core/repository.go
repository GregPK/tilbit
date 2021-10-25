package core

import (
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

func LoadSources() (sources []Source) {
	files, err := ioutil.ReadDir("data/")
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
		err, tilbits = ParseMarkdownBody(string(data))
	} else if strings.Contains(fileString, "My Clippings.txt") {
		err, tilbits = ParseKindleClippingsFile(string(data))
	} else {
		err, tilbits = ParseTextFile(string(data))
	}
	return
}