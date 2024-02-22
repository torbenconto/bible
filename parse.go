package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"unicode"
)

func ParseBible(file *os.File) []Book {
	bible := make([]Book, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line into words
		words := strings.Fields(line)

		// Identify the book name and verse
		bookName := ""
		verseStartIndex := 0
		for i, word := range words {
			if unicode.IsDigit(rune(word[0])) {
				verseStartIndex = i
				break
			}
			bookName += word + " "
		}
		bookName = strings.TrimSpace(bookName)

		// The rest of the line is the verse
		verse := strings.Join(words[verseStartIndex:], " ")

		// Check if the book already exists, if not, create a new book
		var currentBook *Book
		for i := range bible {
			if bible[i].Name == bookName {
				currentBook = &bible[i]
				break
			}
		}
		if currentBook == nil {
			newBook := NewBook(bookName, []Verse{})
			bible = append(bible, *newBook)
			currentBook = &bible[len(bible)-1]
		}

		// Add the verse to the current book
		currentBook.Verses = append(currentBook.Verses, *NewVerse(bookName, verse))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return bible
}
