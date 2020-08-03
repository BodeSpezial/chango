package cmd

import (
	"chango/internal/changes"
	"chango/internal/markdown"
    "chango/internal/files"
	"github.com/spf13/cobra"
	"bufio"
	"log"
	"os"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates the current changelog",
	Long: `Reads all the existing changelog files and bundles them into a changelog with the current day as the releaseday

After bundling the changelog and appending it to the CHANGELOG.md file (or creating a new one if none is existing) it deletes all of the current changelog entries to start a new release cycle.`,
	Run: func(cmd *cobra.Command, args []string) {
		changes := changes.GatherChanges("")
		md := markdown.GenerateMarkdown(changes)

		file, err := os.Open("./Changelog.md")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		files.InsertStringToFile("./Changelog.md", "<!-- begin:changelog -->\n\n"+md, "<!-- begin:changelog -->")

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}