package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/torbenconto/bible"
	"github.com/torbenconto/bible-cli/config"
	"github.com/torbenconto/bible-cli/util"
	"github.com/torbenconto/bible/versions"
	"log"
	"os"
	"strings"
)

var BibleVersion string

func init() {
	cobra.OnInitialize(config.InitDotBible)
	rootCmd.PersistentFlags().StringVarP(&BibleVersion, "version", "v", "NIV", "Specify the version of the Bible to use")
}

var rootCmd = &cobra.Command{
	Use:   "bible [command] [flags]",
	Short: "bible is a command line tool for reading and searching the Bible",
	Long:  `bible is a command line tool for reading and searching the Bible. It supports multiple versions of the Bible and provides a simple interface for reading and searching.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		BibleVersion = strings.ToUpper(BibleVersion)
		if _, ok := versions.VersionMap[BibleVersion]; !ok {
			log.Printf("Version %s not found\n", BibleVersion)
			os.Exit(1)
		}

		newBible := bible.NewBible(versions.VersionMap[BibleVersion])
		util.LoadSourceFile(newBible)

		// Load config file
		apiKey := config.GetApiKey()

		// Create a new context with the API key and the Bible instance as values
		ctx := context.WithValue(context.Background(), "api_key", apiKey)
		ctx = context.WithValue(ctx, "bible", newBible)

		// Set the context to the command
		cmd.SetContext(ctx)
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Call help command
		cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
