package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"fmt"
)

var (
	configCmd = &cobra.Command{
		Use:   "config",
		Short: "View or change your configuration settings for chango.",
		Long: `

To create a config file with the default values run chango config init [flags]`,
		Run: func(cmd *cobra.Command, args []string) {
			if cfgFile := viper.ConfigFileUsed(); cfgFile != "" {
				fmt.Println("Using config file:", cfgFile)
			} else {
				fmt.Println("No config file used yet.\n")
				cmd.Help()
			}
		},
	}

	setCmd = &cobra.Command{
		Use:   "set",
		Short: "Set the given field",
		Run: func(cmd *cobra.Command, args []string) {
			viper.Set(args[0], args[1])
			viper.WriteConfig()
		},
		Args: cobra.ExactArgs(2),
	}

)

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(setCmd)
}
