/*
Copyright © 2021 Emma Nuel

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// cfgCmd represents the cfg command
var cfgCmd = &cobra.Command{
	Use:   "cfg",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cfg called")
		greeting := "Hello"
		name, _ := cmd.Flags().GetString("name") // Check if flag-option is set on the command line
		if name == "" {
			name = "World"
		}
		if viper.GetString("name") != "" { // If variable is defined in configuration file overwrite command line
			// name = viper.GetString("name")
			name = viper.GetString("level1.level2.name")
		}
		if viper.GetString("greeting") != "" { // Check for another option in the
			greeting = viper.GetString("greeting")
		}
		fmt.Println(greeting + " " + name)
	},
}

func init() {
	rootCmd.AddCommand(cfgCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cfgCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cfgCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cfgCmd.Flags().StringP("name", "n", "", "Set name")
}
