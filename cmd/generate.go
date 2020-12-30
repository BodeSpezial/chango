package cmd

import (
	"chango/internal/changes"
	"chango/internal/files"
	"chango/internal/markdown"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"bufio"
	"io/ioutil"
	"log"
	"os"
)

var newChangelogString = `# Changelog

All noteable changes to this project will be documented in this file.

The format is based on [Keep a Changelog]

<!-- begin:changelog -->

<!-- Links -->

[Keep a Changelog]: https://keepachangelog.com/

<!-- Versions -->

[unreleased]: ` + viper.GetString("repoUrl") + `/compare/trunk...develop
[released]: ` + viper.GetString("repoUrl") + `/releases`

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates the current changelog",
	Long: `Reads all the existing changelog files and bundles them into a changelog with the current day as the releaseday

After bundling the changelog and appending it to the CHANGELOG.md file (or creating a new one if none is existing) it deletes all of the current changelog entries to start a new release cycle.`,
	Run: func(cmd *cobra.Command, args []string) {
		changes := changes.GatherChanges("")
		md := markdown.GenerateMarkdown(changes)
		changelogFile := "./" + viper.GetString("changelog.filename") + "." + viper.GetString("changelog.filetype")

		file, err := os.Open(changelogFile)
		if err != nil {
			if _, ok := err.(*os.PathError); ok {
				file, err = os.Create(changelogFile)
				ioutil.WriteFile(changelogFile, []byte(newChangelogString), 0644)
				files.InsertStringToFile(changelogFile, "<!-- begin:changelog -->\n\n"+md, "<!-- begin:changelog -->")
				return
			}
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
        //TODO Filter for date or release according to which schedule parameter is set
		files.InsertStringToFile(changelogFile, "<!-- begin:changelog -->\n\n"+md, "<!-- begin:changelog -->")

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
