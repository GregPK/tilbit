package cmd

import (
	"github.com/spf13/cobra"
)

var (
	addCmd = &cobra.Command{
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
)
