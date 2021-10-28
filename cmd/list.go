package cmd

import (
	"fmt"

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

			data, err := yaml.Marshal(&tilbits)

			if err != nil {
				fmt.Printf("err: %v\n", err)
				return
			}
			fmt.Println(string(data))
			// fmt.Println(fmt.Sprintf("Data lenght: %s", strconv.Itoa(len(string(data)))))
		},
	}
)
