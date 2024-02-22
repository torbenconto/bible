package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
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
	InitializeBible()
}

func main() {

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

	fmt.Println(ParseBible(file))
}
