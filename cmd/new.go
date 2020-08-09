package cmd

import (
	"chango/internal/git"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"log"
	"os"
	"os/exec"
)

func chlogFileName() string {
	return "./" + viper.GetString("path_to_changelogs") + "/" + git.GetCurrentBranch() + ".yml"
}

func openChlogFile() {
	changelogFile := chlogFileName()

	openChlog := exec.Command(viper.GetString("editor"), changelogFile)
	openChlog.Stdin = os.Stdin
	openChlog.Stdout = os.Stdout
	if err := openChlog.Run(); err != nil {
		log.Println(err)
	}

	if viper.GetBool("git.auto_commit") == true {
		git.CommitChangelog(changelogFile)
	}
}

func addChlogEntryToFile() {
	changelogFile := chlogFileName()
	os.Create(changelogFile)
}

var (
	newChlog = &cobra.Command{
		Use:   "new",
		Short: "Add a new changelog entry",
		Long: `Calling this command lets you add a new changelog entry. You can to this directly through the command-line or open an editor.

If using an editor you can either give a name or a path through the -E flag or it uses the one specified in your config-file.`,
		Run: func(cmd *cobra.Command, args []string) {
			addChlogEntryToFile()
			openChlogFile()
		},
	}
	//TODO: Append add, remove, change, fix to file
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
