package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/torbenconto/bible/versions"
	"log"
	"os"
	"path/filepath"
)

var versionsCmd = &cobra.Command{
	Use:   "versions",
	Short: "List all available Bible versions",
	Run: func(cmd *cobra.Command, args []string) {
		// Get home directory
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatal("Error getting user home directory")
		}

		for _, version := range versions.Versions {
			var installedLocally bool = false

			versionPath := filepath.Join(home, ".bible", "versions", version.Name, fmt.Sprintf("%s.txt", version.Name))

			// Check if the version is installed locally
			if _, err := os.Stat(versionPath); err == nil {
				installedLocally = true
			}

			if installedLocally {
				fmt.Printf("%s (installed)\n", version.Name)
			} else {
				fmt.Printf("%s\n", version.Name)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(versionsCmd)
}
