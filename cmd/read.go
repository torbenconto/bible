package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/torbenconto/bible/bible"
)

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read the Bible",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get the Bible from the context
		ctxBible := bible.GetFromContext(cmd.Context())

		reader := bufio.NewReader(os.Stdin)

		startVerse := args[0]

		startReading := false

		for _, book := range ctxBible.Books {
			for _, verse := range book.Verses {
				// Remove book name from verse name
				if strings.ToLower(verse.Name) == strings.ToLower(startVerse) {
					fmt.Println("Press enter to show the next verse or q and then enter to quit")
					startReading = true
				}

				if startReading {
					name := strings.Replace(verse.Name, book.Name+" ", "", 1)
					fmt.Printf("%s: %s\n", name, verse.Text)
					input, _ := reader.ReadString('\n')
					if input == "q\n" {
						return
					}
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
}
