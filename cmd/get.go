package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/torbenconto/bible/bible"
	"log"
	"strings"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a verse from the Bible",
	Long:  `Get a verse from the Bible`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		// Get the Bible from the context
		bible := bible.GetFromContext(cmd.Context())

		targetVerse := args[0]

		splitVerse := strings.Split(targetVerse, " ")

		verses := bible.ParseVerse(splitVerse)

		if len(verses) == 0 {
			log.Fatal("No results found")
		}

		for _, v := range verses {
			fmt.Println(v.Name, v.Text)
		}

	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
