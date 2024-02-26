package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/torbenconto/bible/bible"
	"github.com/torbenconto/bible/config"
	"github.com/torbenconto/bible/versions"
	"log"
	"os"
)

var BibleVersion string

func init() {
	cobra.OnInitialize(config.InitDotBible)
	rootCmd.PersistentFlags().StringVarP(&BibleVersion, "version", "v", "KJV", "Specify the version of the Bible to use")
}

var rootCmd = &cobra.Command{
	Use:   "bible [command] [flags]",
	Short: "bible is a command line tool for reading and searching the Bible",
	Long:  `bible is a command line tool for reading and searching the Bible. It supports multiple versions of the Bible and provides a simple interface for reading and searching.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if _, ok := versions.VersionMap[BibleVersion]; !ok {
			log.Printf("Version %s not found\n", BibleVersion)
			os.Exit(1)
		}

		newBible := bible.NewBible(versions.VersionMap[BibleVersion])
		newBible.LoadSourceFile()

		cmd.SetContext(context.WithValue(context.Background(), "bible", newBible))
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
