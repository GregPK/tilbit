package cmd

import (
	"fmt"
	"os"

	"github.com/GregPK/tilbit/core"
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

func addTil(content string, source string) {
	fmt.Printf("Adding [%s] with source [%s]\n", content, source)
	f, err := os.OpenFile(privateDbFilename(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	tilLine := core.MakeTilLine(content, source)

	if _, err = f.WriteString(tilLine); err != nil {
		panic(err)
	}
}
