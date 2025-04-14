package cmd

import (
	"github.com/damian-elenbaas/wp-plugin-generator/internal/generator"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate [company/project-name]",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Short: "Generate the WordPress plugin boilercode.",
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		err := generator.GeneratePlugin(exportDir, projectName)
		if err != nil {
			cmd.PrintErrf("Error generating plugin: %v\n", err)
			return
		}
	},
}
