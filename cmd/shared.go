package cmd

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/GregPK/tilbit/core"
)

func dbDir() string {
	homeDir, _ := os.UserHomeDir()
	return homeDir + "/.config/tilbit/data/"
}

func privateDbFilename() string {
	return dbDir() + "private.txt"
}

func parseFile(fileString string) (tilbits []core.Tilbit) {
	fileParts := strings.Split(fileString, ".")
	extension := fileParts[len(fileParts)-1]

	if extension == "md" {
		data, err := ioutil.ReadFile(fileString)
		if err != nil {
			panic(err)
		}

		err, tilbits = core.ParseMarkdownBody(string(data))
	} else {
		tilbits = parseTextFile(fileString)
	}
	return
}

// TODO: move to parsers
func parseTextFile(file string) (tilbits []core.Tilbit) {

	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " ")

		if line == "" {
			continue
		}

		parts := strings.Split(line, "{")
		if len(parts) != 2 {
			panic(errors.New(fmt.Sprintf("Unexpected parse of line with metadata: %s", line)))
		}

		text := parts[0]
		jsonstr := "{" + parts[1]

		var tilbit core.Tilbit
		var metadata core.TilbitData
		json.Unmarshal([]byte(jsonstr), &metadata)
		tilbit.Text = text
		tilbit.Data = metadata

		tilbits = append(tilbits, tilbit)
	}
	return
}
