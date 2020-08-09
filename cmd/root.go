package cmd

import (
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"fmt"
	"os"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "chango",
	Short: "A brief description of your application",
	Long: `With chango you can keep and create a well styled changelog as fast as you can type.

Just add your changelog messages along the way with the below described commands and bundle them into one Changelog.md file when it's time for your next release.
The style of the changelog follows the rules of Keep-A-Changelog (read more: https://keep-a-changelog.com)`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// set the default config values
	viper.SetDefault("editor", "vim")
	viper.SetDefault("git.auto_commit", false)
	viper.SetDefault("git.repo_url", "")
	viper.SetDefault("release.schedule", "daily")
	viper.SetDefault("path_to_changelogs", "changelog")
	viper.SetDefault("changelog.filename", "Changelog")
	viper.SetDefault("changelog.filetype", "md")

	// configure the config file
	viper.SetConfigName("chango")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./")
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Search config in home directory with name "chango" (without extension).
	viper.AddConfigPath(home + ".config/chango/")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {

			fmt.Println(`Sorry! It seems that I had trouble reading your config file.
If this keeps happening please contact my maintainer over GitHub: https://github.com/BodeSpezial/chango/issues`)
		}
	}
}
