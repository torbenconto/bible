package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/torbenconto/bible/versions"
)

var versionsCmd = &cobra.Command{
	Use:   "versions",
	Short: "List all available Bible versions",
	Run: func(cmd *cobra.Command, args []string) {
		for version := range versions.VersionMap {
			fmt.Println(version)
		}
	},
}

func init() {
	rootCmd.AddCommand(versionsCmd)
}
