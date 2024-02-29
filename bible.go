package bible

import (
	"bufio"
	"github.com/torbenconto/bible/versions"
	"os"
	"strconv"
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

func (b *Bible) LoadSourceFile(file *os.File) error {
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

		verseChapter := strings.Split(strings.Replace(verseName, bookName, "", 1), ":")[0]

		verseChapterInt, err := strconv.Atoi(strings.Trim(verseChapter, " "))
		if err != nil {
			return err
		}

		// Check if the book already exists, if not, create a new book
		var currentBook *Book
		for i := range b.Books {
			if b.Books[i].Name == bookName {
				currentBook = &b.Books[i]
				break
			}
		}
		if currentBook == nil {
			newBook := NewBook(bookName, []Chapter{})
			b.Books = append(b.Books, *newBook)
			currentBook = &b.Books[len(b.Books)-1]
		}

		newVerse := NewVerse(verseName, verseText)

		// Check if the chapter already exists, if not, create a new chapter
		var currentChapter *Chapter
		for i := range currentBook.Chapters {
			if currentBook.Chapters[i].Number == verseChapterInt {
				currentChapter = &currentBook.Chapters[i]
				break
			}
		}
		if currentChapter == nil {
			newChapter := NewChapter(verseChapterInt, []Verse{*newVerse})
			currentBook.Chapters = append(currentBook.Chapters, *newChapter)
			currentChapter = &currentBook.Chapters[len(currentBook.Chapters)-1]
		} else {
			// Add the verse to the current chapter
			currentChapter.Verses = append(currentChapter.Verses, *newVerse)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func (b *Bible) GetVerse(verse string) []Verse {
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
		endFound := false
		if endFound {
			break
		}

		bookName = strings.ToLower(bookName)

		if strings.ToLower(book.Name) == bookName {
			// Find the verses
			startFound := false

			for _, c := range book.Chapters {
				if endFound {
					break
				}
				for _, v := range c.Verses {
					if strings.ToLower(v.Name) == bookName+" "+verseStart {
						startFound = true
					}
					if startFound {
						verses = append(verses, v)
					}
					if strings.ToLower(v.Name) == bookName+" "+verseEnd {
						endFound = true
						break
					}
				}
			}
		}
	}

	return verses
}
