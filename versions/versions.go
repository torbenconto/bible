package versions

type Version struct {
	Name     string
	Path     string
	Url      string
	BadLines []string
}

func NewVersion(name, path, url string, badLines []string) *Version {
	return &Version{name, path, url, badLines}
}

var KJV = *NewVersion("kjv", "kjv.txt", "https://openbible.com/textfiles/kjv.txt", []string{"KJV", "King James Bible: Pure Cambridge Edition - Text courtesy of www.BibleProtector.com"})
var ASV = *NewVersion("asv", "asv.txt", "https://openbible.com/textfiles/asv.txt", []string{"ASV", "American Standard Version"})

var Versions = []Version{
	KJV,
	ASV,
}

var VersionMap = map[string]Version{
	"KJV":                       KJV,
	"King James Version":        KJV,
	"ASV":                       ASV,
	"American Standard Version": ASV,
}
