package bible

type Book struct {
	Name     string
	Chapters []Chapter
}

func NewBook(name string, chapters []Chapter) *Book {
	return &Book{name, chapters}
}
