package bible

type Verse struct {
	Name string
	Text string
}

func NewVerse(name, text string) *Verse {
	return &Verse{name, text}
}
