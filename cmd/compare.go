package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/torbenconto/bible"
	"github.com/torbenconto/bible-cli/config"
	"github.com/torbenconto/bible-cli/util"
	"github.com/torbenconto/bible/versions"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var compareCmd = &cobra.Command{
	Use:   "compare [verse] [version1] [version2] ... [versionN]",
	Short: "Compares a verse in different bible versions",
	Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		verse := args[0]
		targetVersions := args[1:]

		for _, version := range targetVersions {
			version = strings.ToUpper(version)
		}

		// Get the Bible from the context
		ctxBible := util.GetFromContext(cmd.Context())

		var targetBible = ctxBible

		for _, version := range targetVersions {
			if _, ok := versions.VersionMap[version]; !ok {
				log.Fatalf("Version %s not found", version)
			}

			if targetBible.Version.Name != versions.VersionMap[version].Name {
				newBible := bible.NewBible(versions.VersionMap[version])

				home, err := os.UserHomeDir()
				if err != nil {
					log.Fatal(err)
				}
				file, err := os.Open(filepath.Join(home, fmt.Sprintf(".bible/versions/%s/%s.txt", newBible.Version.Name, newBible.Version.Name)))
				if err != nil {
					if os.IsNotExist(err) {
						log.Printf("Version %s not found locally", newBible.Version.Name)
						log.Println("Downloading the version")
						config.InitVersion(newBible.Version)

						// Bad but only way to make it look clean
						os.Exit(1)
					}
				}

				err = newBible.LoadSourceFile(file)
				if err != nil {
					log.Fatal(err)
				}
				targetBible = newBible
			}
		}

		for _, version := range targetVersions {
			verses := targetBible.GetVerse(verse)

			if len(verses) == 0 {
				log.Fatal("No results found")
			}

			for _, verse := range verses {
				fmt.Println(verse.Name, verse.Text, "|", version)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(compareCmd)
}
