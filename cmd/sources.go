package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
)

var (
	sourcesCmd = &cobra.Command{
		Use:   "sources",
		Short: "Print information about tilbit sources",
		Long:  `Get debug info about all sources of TILBit items.`,
		// AddArgument("content", "Body of the TILBit", "").
		// AddArgument("source", "Source of the TILBit", "").
		Run: func(cmd *cobra.Command, args []string) {
			printStats()
		},
	}
)

func printStats() {
	files, err := ioutil.ReadDir("data/")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		name := dbDir() + file.Name()

		if !file.IsDir() {
			bits := parseFile(name)
			fmt.Println(name, len(bits))
		}
	}
}
