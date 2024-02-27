package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/torbenconto/bible/config"
)

var configCmd = &cobra.Command{
	Use:   "config [api_key]",
	Short: "Configure AI features of Bible",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Set the API key
		config.InitApiKey(args[0])

		fmt.Println("API key set successfully")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
