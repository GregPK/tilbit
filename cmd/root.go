package cmd

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/GregPK/tilbit/core"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "tilbit",
		Short: "TILBit",
		Long: `Write down your learnings.
					 Revise them on each new terminal window.`,
		Run: func(cmd *cobra.Command, args []string) {
			tilbits := parseFile(privateDbFilename())

			randTil := getRandomBit(tilbits)

			fmt.Println(core.GetBitString(randTil))
		},
	}
)

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(sourcesCmd)
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func getRandomBit(tilbits []core.Tilbit) (randomTilbit core.Tilbit) {
	rand.Seed(time.Now().UnixNano())
	// fmt.Printf("%s:\n", len(tilbits))
	randomTilbit = tilbits[rand.Intn(len(tilbits))]
	return
}
