package cmd

import (
	"github.com/GregPK/tilbit/core"
	"github.com/spf13/cobra"
)

type Configuration struct {
	outputFormat string
}

var Config = Configuration{"box"}

var (
	rootCmd = &cobra.Command{
		Use:   "tilbit",
		Short: "TILBit",
		Long: `Write down your learnings.
					 Revise them on each new terminal window.`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				args = append(args, "random")
			}
			showCmd.Run(cmd, args)
		},
		Version: core.VERSION,
	}
)

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(sourcesCmd)
	rootCmd.AddCommand(showCmd)
	rootCmd.SetVersionTemplate("TILBit version: {{.Version}}\n")
	addFlags()
}

func addFlags() {
	showCmd.Flags().StringVarP(&Config.outputFormat, "output-format", "f", "box", "Output format for show command")
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
