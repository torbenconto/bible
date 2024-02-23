package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"unicode"
)

var versionName string

func init() {
	flag.StringVar(&versionName, "version", "KJV", "Specify the version of the Bible to use")
	flag.Parse()

	InitDotBible()
}

func main() {

	versionMap := map[string]Version{
		"KJV":                       KJV,
		"King James Version":        KJV,
		"ASV":                       ASV,
		"American Standard Version": ASV,
	}

	if _, ok := versionMap[versionName]; !ok {
		log.Fatalf("Version %s not found", versionName)
	}

	var bible = NewBible(versionMap[versionName])

	bible.LoadSourceFile()

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

			// Find the book
			for _, book := range bible.Books {
				if book.Name == bookName {
					// Find the verses
					startFound := false
					for _, v := range book.Verses {
						if v.Name == bookName+" "+verseStart {
							startFound = true
						}
						if startFound {
							fmt.Println(v.Name, v.Text)
						}
						if v.Name == bookName+" "+verseEnd {
							break
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

				for _, b := range bible.Books {
					if b.Name == bookName {
						book = b
						break
					}
				}
			} else {
				book = bible.Books[rand.Intn(len(bible.Books))]
			}

			verse := book.Verses[rand.Intn(len(book.Verses))]

			fmt.Println(verse.Name, verse.Text)

		}
	}
}
