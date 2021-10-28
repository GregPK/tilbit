package cmd

import (
	"github.com/GregPK/tilbit/core"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List TILBits",
		Long:  `Lists all TILBits from all available sources and outputs all data that is has for a given entry`,
		Run: func(cmd *cobra.Command, args []string) {
			tilbits := core.AllTilbits()

			data, _ := yaml.Marshal(&tilbits)
			println(string(data))
		},
	}
)
