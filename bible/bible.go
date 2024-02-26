package bible

import (
	"bufio"
	"fmt"
	"github.com/torbenconto/bible/config"
	"github.com/torbenconto/bible/versions"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

type Bible struct {
	Version versions.Version
	Books   []Book
}

func NewBible(version versions.Version) *Bible {
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
			config.InitVersion(b.Version)
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

func (b *Bible) ParseVerse(verse string) []Verse {
	targetVerse := strings.Split(verse, " ")

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
	splitVerse := strings.Join(targetVerse, " ")

	// Split the verse into start and end
	verseRange := strings.Split(splitVerse, "-")
	verseStart := verseRange[0]
	verseEnd := ""

	verseStartSplit := strings.Split(verseStart, ":")

	if len(verseRange) > 1 {
		verseEnd = verseStartSplit[0] + ":" + verseRange[1]
	} else {
		verseEnd = verseStart
	}

	verses := make([]Verse, 0)

	// Find the book
	for _, book := range b.Books {
		bookName = strings.ToLower(bookName)

		if strings.ToLower(book.Name) == bookName {
			// Find the verses
			startFound := false
			for _, v := range book.Verses {
				if strings.ToLower(v.Name) == bookName+" "+verseStart {
					startFound = true
				}
				if startFound {
					verses = append(verses, v)
				}
				if strings.ToLower(v.Name) == bookName+" "+verseEnd {
					break
				}
			}
		}
	}

	return verses
}

func (b *Bible) Search(query string, caseSensitive bool) []Verse {
	verses := make([]Verse, 0)
	for _, book := range b.Books {
		for _, verse := range book.Verses {
			var text string
			if !caseSensitive {
				query = strings.ToLower(query)
				text = strings.ToLower(verse.Text)
			} else {
				text = verse.Text
			}
			if strings.Contains(text, query) {
				verses = append(verses, verse)
			}
		}
	}
	return verses
}
