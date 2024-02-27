package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/torbenconto/bible-cli/util"
)

var booksCmd = &cobra.Command{
	Use:   "books",
	Short: "List all books in the Bible",
	Run: func(cmd *cobra.Command, args []string) {
		// Get the Bible from the context
		ctxBible := util.GetFromContext(cmd.Context())

		for _, book := range ctxBible.Books {
			fmt.Println(book.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(booksCmd)
}
