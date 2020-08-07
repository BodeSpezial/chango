package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"fmt"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "View or change your configuration settings for chango.",
	Long: `

If you run this command for the first time and you don't have a config file yet it creates a new one with the default values.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				viper.SafeWriteConfig()
			} else {
				panic(fmt.Errorf("Fatal error config file: %s \n", err))
			}
		}
		fmt.Println("View or change your configuration settings for chango.")

	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
