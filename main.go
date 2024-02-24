package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

var versionName string

func init() {
	flag.StringVar(&versionName, "version", "KJV", "Specify the version of the Bible to use")
	flag.Parse()

	InitDotBible()
}

func main() {

	if _, ok := VersionMap[versionName]; !ok {
		log.Fatalf("Version %s not found", versionName)
	}

	var bible = NewBible(VersionMap[versionName])

	bible.LoadSourceFile()

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "get":
			if len(os.Args) < 3 {
				log.Println("Usage: bible get <verse>")
				log.Fatal("Not enough arguments")
			}
			targetVerse := strings.Split(os.Args[2], " ")

			verses := bible.ParseVerse(targetVerse)

			for _, verse := range verses {
				fmt.Println(verse.Name, verse.Text)
			}

		case "random":
			bookName := ""
			if len(os.Args) > 2 {
				bookName = os.Args[2]
			}

			var book Book
			if bookName != "" {
				for _, b := range bible.Books {
					if b.Name == bookName {
						book = b
						break
					}
				}

				if book.Name == "" {
					log.Fatalf("Custom book not found, run bible books to retireve a list of availible books")
				}
			} else {
				book = bible.Books[rand.Intn(len(bible.Books))]
			}

			verse := book.Verses[rand.Intn(len(book.Verses))]

			fmt.Println(verse.Name, verse.Text)
		case "compare":
			if len(os.Args) < 5 {
				log.Println("Usage: bible compare <verse> <version1> <version2> ... <versionN>")
				log.Fatal("Not enough arguments")
			}

			targetVerse := strings.Split(os.Args[2], " ")

			versions := os.Args[3:]

			for _, version := range versions {
				if _, ok := VersionMap[version]; !ok {
					log.Fatalf("Version %s not found", version)
				}

				var bible = NewBible(VersionMap[version])
				bible.LoadSourceFile()

				verses := bible.ParseVerse(targetVerse)

				for _, verse := range verses {
					fmt.Println(verse.Name, verse.Text, "|", version)
				}
			}
		case "search":
			if len(os.Args) < 3 {
				log.Println("Usage: bible search <term>")
				log.Fatal("Not enough arguments")
			}
			term := os.Args[2]
			verses := bible.Search(term, false)

			for _, verse := range verses {
				fmt.Println(verse.Name, verse.Text)
			}
		case "books":
			for _, book := range bible.Books {
				fmt.Println(book.Name)
			}
		default:
			log.Println("Usage: bible <command> [args]")
		}
	}
}
