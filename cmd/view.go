package cmd

import (
	"chango/internal/changes"
	"chango/internal/markdown"
	"fmt"

	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "Prints the Changelog that will be created",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		changes := changes.GatherChanges("")
		fmt.Println(
			markdown.GenerateMarkdown(changes))
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)

}
