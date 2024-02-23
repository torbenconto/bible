package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

type Bible struct {
	Version Version
	Books   []Book
}

func NewBible(version Version) *Bible {
	return &Bible{Version: version}
}

func (b *Bible) LoadSourceFile() *Bible {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(filepath.Join(home, fmt.Sprintf(".bible/versions/%s/%s.txt", b.Version.Name, b.Version.Name)))
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Skip first two lines
		for _, badLine := range b.Version.BadLines {
			if scanner.Text() == badLine {
				continue
			}
		}
		line := scanner.Text()

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
		var currentBook *Book
		for i := range b.Books {
			if b.Books[i].Name == bookName {
				currentBook = &b.Books[i]
				break
			}
		}
		if currentBook == nil {
			newBook := NewBook(bookName, []Verse{})
			b.Books = append(b.Books, *newBook)
			currentBook = &b.Books[len(b.Books)-1]
		}

		// Add the verse to the current book
		currentBook.Verses = append(currentBook.Verses, *NewVerse(verseName, verseText))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return b
}
