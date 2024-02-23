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
		currentBook.Verses = append(currentBook.Verses, *NewVerse(verseName, verseText))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return bible
}
