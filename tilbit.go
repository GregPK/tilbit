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
			f, _ := os.Open("data/database.txt")
			scanner := bufio.NewScanner(f)

			var tilbits = []Tilbit{}

			for scanner.Scan() {
				line := strings.Trim(scanner.Text(), " ")

				if line == "" {
					continue
				}

				// Split the line on commas.
				parts := strings.Split(line, "{")

				if len(parts) != 2 {
					errors.New(fmt.Sprintf("Unexpected parse of line with metadata: %s", line))
				}
				// Loop over the parts from the string.

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
			rand.Seed(time.Now().UnixNano())
			randomIndex := rand.Intn(len(tilbits))

			randTil := tilbits[randomIndex]
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

func getBitString(tilbit Tilbit) (str string, err error) {
	text := wordwrap.WrapString(tilbit.Text, 120)
	footer := fmt.Sprintf("   -- %s (%s)", tilbit.Data.Source, tilbit.Data.AddedOn)
	str = fmt.Sprintf("%s\n%s", text, footer)
	return
}
