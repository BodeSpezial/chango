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

	getCmd = &cobra.Command{
		Use:   "get",
		Short: "Print the value of the given field",
		Run: func(cmd *cobra.Command, args []string) {
			if args[0] == "path" {
				if cfgFile := viper.ConfigFileUsed(); cfgFile != "" {
					fmt.Println("Using config file:", cfgFile)
				} else {
					fmt.Println("No config file used yet.")
				}

			} else {
				fmt.Println(viper.Get(args[0]))
			}
		},
		Args: cobra.ExactArgs(1),
	}

	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initializes a configuration",
		Long:  "Writes a file with all the standard values if none exists.",
		Run: func(cmd *cobra.Command, args []string) {
			viper.SafeWriteConfig()
		},
	}
)

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(setCmd)
	configCmd.AddCommand(getCmd)
	configCmd.AddCommand(initCmd)
}
