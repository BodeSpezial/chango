package cmd

import (
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates the current changelog",
	Long: `Reads all the existing changelog files and bundles them into a changelog with the current day as the releaseday

After bundling the changelog and appending it to the CHANGELOG.md file (or creating a new one if none is existing) it deletes all of the current changelog entries to start a new release cycle.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
