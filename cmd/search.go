package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/torbenconto/bible/bible"
	"log"
	"strings"
)

var searchCmd = &cobra.Command{
	Use:   "search [term]",
	Short: "Search for terms in the Bible",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get the Bible from the context
		ctxBible := bible.GetFromContext(cmd.Context())

		caseSensitive, err := cmd.Flags().GetBool("case-sensitive")
		if err != nil {
			log.Fatalf("Error getting case-sensitive: %s", err)
		}
		bookName := cmd.Flag("book").Value.String()
		maxResults, err := cmd.Flags().GetInt("max")
		if err != nil {
			log.Fatalf("Error getting max: %s", err)
		}
		wordsOnly, err := cmd.Flags().GetBool("words-only")
		if err != nil {
			log.Fatalf("Error getting words-only: %s", err)
		}

		verse := args[0]

		verses := make([]bible.Verse, 0)
		var query = verse
		for _, book := range ctxBible.Books {
			// If bookName is provided, only search in that book
			if bookName != "" && book.Name != bookName {
				continue
			}
			for _, verse := range book.Verses {
				var text string
				if !caseSensitive {
					query = strings.ToLower(query)
					text = strings.ToLower(verse.Text)
				} else {
					text = verse.Text
				}
				if wordsOnly {
					words := strings.Split(text, " ")
					for _, word := range words {
						if word == query {
							verses = append(verses, verse)
							break
						}
					}
				} else {
					if strings.Contains(text, query) {
						verses = append(verses, verse)
					}
				}
			}
		}
		if len(verses) == 0 {
			log.Fatal("No results found")
		}

		if maxResults > 0 && len(verses) > maxResults {
			verses = verses[:maxResults]
		}

		for _, v := range verses {
			fmt.Println(v.Name, v.Text)
		}
	},
}

func init() {
	searchCmd.Flags().BoolP("case-sensitive", "c", false, "Make the search case sensitive")
	searchCmd.Flags().StringP("book", "b", "", "Specify a book to search in")
	searchCmd.Flags().BoolP("words-only", "w", true, "Only return results that match the entire query")
	searchCmd.Flags().IntP("max", "m", -1, "Maximum number of results to return")
	rootCmd.AddCommand(searchCmd)
}
