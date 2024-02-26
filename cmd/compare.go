package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/torbenconto/bible/bible"
	"github.com/torbenconto/bible/versions"
	"log"
)

var compareCmd = &cobra.Command{
	Use:   "compare [verse] [version1] [version2] ... [versionN]",
	Short: "Compares a verse in different bible versions",
	Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		verse := args[0]
		targetVersions := args[1:]

		// Get the Bible from the context
		ctxBible := bible.GetFromContext(cmd.Context())

		for _, version := range targetVersions {
			var targetBible = ctxBible

			if _, ok := versions.VersionMap[version]; !ok {
				log.Fatalf("Version %s not found", version)
			}

			if targetBible.Version.Name != versions.VersionMap[version].Name {
				newBible := bible.NewBible(versions.VersionMap[version])
				newBible.LoadSourceFile()
				targetBible = newBible
			}

			verses := targetBible.ParseVerse(verse)

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
