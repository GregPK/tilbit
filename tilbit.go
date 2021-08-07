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

	"github.com/spf13/cobra"
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
	var rootCmd = &cobra.Command{
		Use:   "tilbit",
		Short: "TILBit",
		Long: `A Fast and Flexible Static Site Generator built with
					  love by spf13 and friends in Go.
					  Complete documentation is available at http://hugo.spf13.com`,
		Run: func(cmd *cobra.Command, args []string) {
			tilbits := parseFile(privateDbFilename())

			randTil := getRandomBit(tilbits)

			fmt.Println(getBitString(randTil))
		},
	  }

	var addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add a TILBit",
		Long:  `This commands add an item to the private TILBit database.`,
		// AddArgument("content", "Body of the TILBit", "").
		// AddArgument("source", "Source of the TILBit", "").
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			addTil(args[0], args[1])
		},
	}
	rootCmd.AddCommand(addCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
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
