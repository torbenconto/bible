package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

type Book struct {
	Name   string
	Verses []Verse
}

func NewBook(name string, verses []Verse) *Book {
	return &Book{name, verses}
}

type Verse struct {
	Name string
	Text string
}

func NewVerse(name, text string) *Verse {
	return &Verse{name, text}
}

func init() {
	InitDotBible()
}

func main() {
	var bible []Book

	// Home dir
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	// Get kjv version for now
	file, err := os.Open(filepath.Join(home, ".bible/versions/kjv/kjv.txt"))
	if err != nil {
		log.Fatal(err)
	}

	bible = ParseBible(file)

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "get":
			targetVerse := strings.Split(os.Args[2], " ")
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

			// Find the book
			for _, book := range bible {
				if book.Name == bookName {
					// Find the verse
					for _, v := range book.Verses {
						if v.Name == bookName+" "+verse {
							fmt.Println(v.Name, v.Text)
							return
						}
					}
				}
			}
		case "random":
			fmt.Print("Do you want to use a custom book? (yes/no) ")
			var customBook string
			_, err := fmt.Scan(&customBook)
			if err != nil {
				log.Fatal(err)
			}

			var book Book
			if strings.ToLower(customBook) == "yes" {
				fmt.Print("Enter the name of the book: ")
				var bookName string
				_, err := fmt.Scan(&bookName)
				if err != nil {
					log.Fatal(err)
				}

				for _, b := range bible {
					if b.Name == bookName {
						book = b
						break
					}
				}
			} else {
				book = bible[rand.Intn(len(bible))]
			}

			verse := book.Verses[rand.Intn(len(book.Verses))]

			fmt.Println(verse.Name, verse.Text)

		}

	}
}
