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
		SetVersion("0.0.2").
		SetDescription("TILBit")

	// configure the root command
	commando.
		Register(nil).
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			tilbits := parseFile(privateDbFilename())

			randTil := getRandomBit(tilbits)

			fmt.Println(getBitString(randTil))
		})

	// configure info command
	commando.
		Register("add").
		SetShortDescription("Add a TILBit").
		SetDescription("This commands add an item to the private TILBit database.").
		AddArgument("content", "Body of the TILBit", "").
		AddArgument("source", "Source of the TILBit", "").
		// AddFlag("stdin,s", "", commando.Int, nil). // required
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			addTil(args["content"].Value, args["source"].Value)
		})

	// parse command-line arguments
	commando.Parse(nil)
}

func privateDbFilename() string {
	homeDir, _ := os.UserHomeDir()
	return homeDir + "/.config/tilbit/data/private.txt"
}

func parseFile(file string) (tilbits []Tilbit) {
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


func addTil(content string, source string) {
	fmt.Printf("Adding [%s] with source [%s]\n", content, source)
	f, err := os.OpenFile(privateDbFilename(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	tilLine := makeTilLine(content, source)

	if _, err = f.WriteString(tilLine); err != nil {
		panic(err)
	}
}

func getRandomBit(tilbits []Tilbit) (randomTilbit Tilbit) {
	rand.Seed(time.Now().UnixNano())
	// fmt.Printf("%s:\n", len(tilbits))
	randomTilbit = tilbits[rand.Intn(len(tilbits))]
	return
}
