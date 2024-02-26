package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/torbenconto/bible/versions"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall a Bible version",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Capitalize the first letter of the version
		version := strings.ToUpper(args[0])

		// Check if valid
		if _, ok := versions.VersionMap[version]; !ok {
			log.Fatalf("Version %s not found", version)
		}

		// Get the user's home directory
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("Error getting home directory: %s", err)
		}

		// Construct the path to the version directory
		versionDir := filepath.Join(home, ".bible", "versions", strings.ToLower(version))

		if _, err := os.Stat(versionDir); os.IsNotExist(err) {
			log.Fatalf("Version %s is not installed", version)
		}

		// Delete the version directory
		err = os.RemoveAll(versionDir)
		if err != nil {
			log.Fatalf("Error uninstalling version: %s", err)
		}

		fmt.Printf("Version %s uninstalled successfully\n", version)
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
