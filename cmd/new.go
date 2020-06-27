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
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Add a new changelog entry",
	Long: `Calling this command lets you add a new changelog entry. You can to this directly through the command-line or open an editor.

If using an editor you can either give a name or a path through the -E flag or it uses the one specified in your config-file.`,
	Run: func(cmd *cobra.Command, args []string) {
		changelogFile := "./" + conf.ChlogFolder + "/" + git.GetCurrentBranch() + ".toml"

		openChlog := exec.Command(conf.Editor, changelogFile)
		openChlog.Stdin = os.Stdin
		openChlog.Stdout = os.Stdout
		if err := openChlog.Run(); err != nil {
			log.Println(err)
		}

		if conf.AutoCommit == true {
			git.CommitChangelog(changelogFile)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
