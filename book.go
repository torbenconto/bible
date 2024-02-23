package main

type Book struct {
	Name   string
	Verses []Verse
}

func NewBook(name string, verses []Verse) *Book {
	return &Book{name, verses}
}
