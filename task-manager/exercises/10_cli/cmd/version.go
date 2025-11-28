package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of taskctl",
	Long:  `All software has versions. This is taskctl`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("taskctl version 0.0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}