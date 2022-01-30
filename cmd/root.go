package cmd

import (
	"github.com/GregPK/tilbit/core"
	"github.com/spf13/cobra"
)

type Configuration struct {
	outputFormat string
	sources      []string
}

var Config = Configuration{"box", make([]string, 0)}

func RootCmd(inputTilbits ...core.Tilbit) *cobra.Command {
	rcmd := cobra.Command{
		Use:   "tilbit",
		Short: "TILBit",
		Long: `Write down your learnings.
					 Revise them on each new terminal window.`,
		Run: func(cmd *cobra.Command, args []string) {
			showCmd := ShowCmd(inputTilbits...)
			if len(args) == 0 {
				args = append(args, "random")
			}
			showCmd.Run(cmd, args)
		},
		Version: core.VERSION,
	}
	return &rcmd
}

var rootCmd = RootCmd()

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(sourcesCmd)
	rootCmd.AddCommand(showCmd)
	rootCmd.SetVersionTemplate("TILBit version: {{.Version}}\n")
	rootCmd.PersistentFlags().StringVarP(&Config.outputFormat, "output-format", "f", "box", "Output format for show command")
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
