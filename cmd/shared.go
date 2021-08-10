package cmd

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/GregPK/tilbit/core"
)

func privateDbFilename() string {
	homeDir, _ := os.UserHomeDir()
	return homeDir + "/.config/tilbit/data/private.txt"
}

func parseFile(file string) (tilbits []core.Tilbit) {
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
			errors.New(fmt.Sprintf("Unexpected parse of line with metadata: %s", line))
		}

		text := parts[0]
		jsonstr := "{" + parts[1]

		var tilbit core.Tilbit
		var metadata core.TilbitData
		json.Unmarshal([]byte(jsonstr), &metadata)
		tilbit.Text = text
		tilbit.Data = metadata

		tilbits = append(tilbits, tilbit)
		// fmt.Println(tilbits[0].Text)
	}
	return
}
