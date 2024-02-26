package cmd

import (
	"github.com/spf13/cobra"
	"github.com/torbenconto/bible/config"
	"github.com/torbenconto/bible/versions"
	"log"
	"strings"
)

var downloadCmd = &cobra.Command{
	Use:   "install [version]",
	Short: "Download a Bible version",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		version := strings.ToUpper(args[0])

		// Check if valid
		if _, ok := versions.VersionMap[version]; !ok {
			log.Fatalf("Version %s not found", version)
		}

		config.InitVersion(versions.VersionMap[version])
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}
