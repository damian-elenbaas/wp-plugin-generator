package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of WP Plugin Generator",
	Long:  `All software has versions. This is WP Plugin Generator's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("WP Plugin Generator v0.1.0 -- HEAD")
	},
}
