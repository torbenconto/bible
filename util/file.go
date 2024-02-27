package util

import (
	"bufio"
	"fmt"
	"github.com/torbenconto/bible"
	"github.com/torbenconto/bible-cli/config"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

func LoadSourceFile(b *bible.Bible) *bible.Bible {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(filepath.Join(home, fmt.Sprintf(".bible/versions/%s/%s.txt", b.Version.Name, b.Version.Name)))
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("Version %s not found locally", b.Version.Name)
			log.Println("Downloading the version")
			config.InitVersion(b.Version)

			// Bad but only way to make it look clean
			os.Exit(1)
		}
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		badline := false

		for _, badLine := range b.Version.BadLines {
			if strings.Contains(line, badLine) {
				badline = true
			}
		}

		if badline {
			continue
		}

		// Split the line into words
		words := strings.Fields(line)

		// Identify the book name and verse
		bookName := ""
		verseStartIndex := 0
		for i, word := range words {
			if unicode.IsDigit(rune(word[0])) && i != 0 { // Ignore the first word starting with a digit
				verseStartIndex = i
				break
			}
			bookName += word + " "
		}

		bookName = strings.TrimSpace(bookName)

		verseName := bookName + " " + words[verseStartIndex]
		verseText := strings.Join(words[verseStartIndex+1:], " ")

		// Check if the book already exists, if not, create a new book
		var currentBook *bible.Book
		for i := range b.Books {
			if b.Books[i].Name == bookName {
				currentBook = &b.Books[i]
				break
			}
		}
		if currentBook == nil {
			newBook := bible.NewBook(bookName, []bible.Verse{})
			b.Books = append(b.Books, *newBook)
			currentBook = &b.Books[len(b.Books)-1]
		}

		// Add the verse to the current book
		currentBook.Verses = append(currentBook.Verses, *bible.NewVerse(verseName, verseText))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return b
}
