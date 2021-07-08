package main

import (
	"fmt"
	// "io/ioutil"
	// "log"
	"bufio"
	"encoding/json"
	"errors"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/mitchellh/go-wordwrap"
	"github.com/thatisuday/commando"
)

type Tilbit struct {
	Text string
	Data TilbitData
}

type TilbitData struct {
	AddedOn string
	Source  string
	Private string
}

func main() {
	// configure commando
	commando.
		SetExecutableName("tilbit").
		SetVersion("0.0.1").
		SetDescription("TIL bit")

	// configure the root command
	commando.
		Register(nil).
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			tilbits := parseFile("data/private.txt")

			randTil := getRandomBit(tilbits)

			fmt.Println(getBitString(randTil))
		})

	// configure info command
	commando.
		Register("info").
		SetShortDescription("displays detailed information of a directory").
		SetDescription("This command displays more information about the contents of the directory like size, permission and ownership, etc.").
		AddArgument("dir", "local directory path", "./").                  // default `./`
		AddFlag("level,l", "level of depth to travel", commando.Int, nil). // required
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			fmt.Printf("Printing options of the `info` command...\n\n")

			// print arguments
			for k, v := range args {
				fmt.Printf("arg -> %v: %v(%T)\n", k, v.Value, v.Value)
			}

			// print flags
			for k, v := range flags {
				fmt.Printf("flag -> %v: %v(%T)\n", k, v.Value, v.Value)
			}
		})

	// parse command-line arguments
	commando.Parse(nil)
}

func parseFile(file string) (tilbits []Tilbit) {
	f, _ := os.Open(file)
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

		var tilbit Tilbit
		var metadata TilbitData
		json.Unmarshal([]byte(jsonstr), &metadata)
		tilbit.Text = text
		tilbit.Data = metadata

		tilbits = append(tilbits, tilbit)
		// fmt.Println(tilbits[0].Text)
	}
	return
}

func getRandomBit(tilbits []Tilbit) (randomTilbit Tilbit) {
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("%s:\n", len(tilbits))
	randomTilbit = tilbits[rand.Intn(len(tilbits))]
	return
}

func getBitString(tilbit Tilbit) (str string, err error) {
	text := wordwrap.WrapString(tilbit.Text, 120)
	footer := fmt.Sprintf("   -- %s (%s)", tilbit.Data.Source, tilbit.Data.AddedOn)
	str = fmt.Sprintf("%s\n%s", text, footer)
	return
}
