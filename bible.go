package bible

import (
	"github.com/torbenconto/bible/versions"
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
