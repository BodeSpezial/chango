package cmd

import (
	"chango/internal/config"
	"chango/internal/git"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

var conf = config.ParseConfigFile()

func chlogFileName() string {
	return "./" + conf.Git.ChlogFolder + "/" + git.GetCurrentBranch() + ".toml"
}

		if conf.AutoCommit == true {
			git.CommitChangelog(changelogFile)
		}
	},
}

var (
	newChlog = &cobra.Command{
		Use:   "new",
		Short: "Add a new changelog entry",
		Long: `Calling this command lets you add a new changelog entry. You can to this directly through the command-line or open an editor.

If using an editor you can either give a name or a path through the -E flag or it uses the one specified in your config-file.`,
		Run: func(cmd *cobra.Command, args []string) {
			openChlogFile()
		},
	}
	add = &cobra.Command{
		Use:   "add",
		Short: "Creates a new changelog entry from type 'Added'",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			addChlogEntryToFile()

			//openChlogFile()
		}}
	remove = &cobra.Command{
		Use:   "remove",
		Short: "Creates a new changelog entry from type 'Removed'",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			addChlogEntryToFile()

			//openChlogFile()
		}}
	change = &cobra.Command{
		Use:   "change",
		Short: "Creates a new changelog entry from type 'Changed'",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			addChlogEntryToFile()

			//openChlogFile()
		}}
	fix = &cobra.Command{
		Use:   "fix",
		Short: "Creates a new changelog entry from type 'Fixed'",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			addChlogEntryToFile()

			//openChlogFile()
		}}
)

func init() {
	rootCmd.AddCommand(newChlog)
	newChlog.AddCommand(add)
	newChlog.AddCommand(change)
	newChlog.AddCommand(fix)
	newChlog.AddCommand(remove)
}
