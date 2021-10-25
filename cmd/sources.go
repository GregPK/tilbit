package cmd

import (
	"log"
	"time"

	"github.com/GregPK/tilbit/core"
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
	sources := core.LoadSources()

	start := time.Now()
	log.Println("Loading sources from " + core.FileRepositoryDir())
	for _, source := range sources {
		println(source.Uri, len(source.Tilbits))
	}
	log.Println("Finished, loaded in: ", time.Since(start))
}
