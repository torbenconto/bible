package bible

type Chapter struct {
	Number int
	Verses []Verse
}

func NewChapter(number int, verses []Verse) *Chapter {
	return &Chapter{number, verses}
}
