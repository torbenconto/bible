package main

type Verse struct {
	Name string
	Text string
}

func NewVerse(name, text string) *Verse {
	return &Verse{name, text}
}

// 1:1-5
