package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	cfgFile   string
	exportDir string

	rootCmd = &cobra.Command{
		Use:   "wp-plugin-generator",
		Short: "A WordPress plugin generator",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&exportDir, "export-dir", ".", "Export directory for the generated plugin")
}

func main() {
	rootCmd.Execute()
}
