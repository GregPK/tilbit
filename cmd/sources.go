package cmd

import (
	"github.com/spf13/cobra"
)

var (
	sourcesCmd = &cobra.Command{
		Use:   "sources",
		Short: "Print information about tilbit sources",
		Long:  `Get debug info about all sources of TILBit items.`,
		Run: func(cmd *cobra.Command, args []string) {
			// No-op - we are moving away from the concept of sources
			return
		},
	}
)
