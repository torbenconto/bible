package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/torbenconto/bible/bible"
	"log"
	"math/rand"
)

var count int
var bookName string

var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Get a random verse from the Bible",
	Run: func(cmd *cobra.Command, args []string) {
		// Get the Bible from the context
		ctxBible := bible.GetFromContext(cmd.Context())

		var book bible.Book
		if bookName != "" {
			for _, b := range ctxBible.Books {
				if b.Name == bookName {
					book = b
					break
				}
			}

			if book.Name == "" {
				log.Fatalf("Custom book not found, run bible books to retireve a list of availible books")
			}
		} else {
			book = ctxBible.Books[rand.Intn(len(ctxBible.Books))]
		}

		for i := 0; i < count; i++ {
			verse := book.Verses[rand.Intn(len(book.Verses))]
			fmt.Println(verse.Name, verse.Text)
		}
	},
}

func init() {
	randomCmd.Flags().IntVarP(&count, "count", "c", 1, "Number of random verses to get")
	randomCmd.Flags().StringVarP(&bookName, "book", "b", "", "Specify a book to get random verses from")
	rootCmd.AddCommand(randomCmd)
}
