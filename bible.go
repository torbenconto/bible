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
		if os.IsNotExist(err) {
			log.Printf("Version %s not found locally", b.Version.Name)
			log.Println("Downloading the version")
			InitVersion(b.Version)
		}
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

func (b *Bible) ParseVerse(targetVerse []string) []*Verse {
	// Get the book
	bookName := ""
	// If the first character is a digit, then the book name is two words
	if unicode.IsDigit(rune(targetVerse[0][0])) {
		bookName = strings.Join(targetVerse[:2], " ")
		targetVerse = targetVerse[2:]
	} else {
		bookName = targetVerse[0]
		targetVerse = targetVerse[1:]
	}

	// Get the verse
	verse := strings.Join(targetVerse, " ")

	// Split the verse into start and end
	verseRange := strings.Split(verse, "-")
	verseStart := verseRange[0]
	verseEnd := ""

	verseStartSplit := strings.Split(verseStart, ":")

	if len(verseRange) > 1 {
		verseEnd = verseStartSplit[0] + ":" + verseRange[1]
	} else {
		verseEnd = verseStart
	}

	verses := make([]*Verse, 0)

	// Find the book
	for _, book := range b.Books {
		if book.Name == bookName {
			// Find the verses
			startFound := false
			for _, v := range book.Verses {
				if v.Name == bookName+" "+verseStart {
					startFound = true
				}
				if startFound {
					verses = append(verses, &v)
				}
				if v.Name == bookName+" "+verseEnd {
					break
				}
			}
		}
	}

	return verses
}
