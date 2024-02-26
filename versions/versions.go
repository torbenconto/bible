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
var BSB = *NewVersion("bsb", "bsb.txt", "https://bereanbible.com/bsb.txt", []string{"The Holy Bible, Berean Standard Bible, BSB is produced in cooperation with Bible Hub, Discovery Bible, OpenBible.com, and the Berean Bible Translation Committee.", "This text of God's Word has been dedicated to the public domain. Free resources and databases are available at BereanBible.com.\t", "Verse\tBerean Standard Bible"})
var DBT = *NewVersion("dbt", "dbt.txt", "https://openbible.com/textfiles/dbt.txt", []string{"DBT", "Darby Bible Translation"})

var Versions = []Version{
	KJV,
	ASV,
	BSB,
	DBT,
}

var VersionMap = map[string]Version{
	"KJV":                       KJV,
	"King James Version":        KJV,
	"ASV":                       ASV,
	"American Standard Version": ASV,
	"BSB":                       BSB,
	"Berean Standard Bible":     BSB,
	"DBT":                       DBT,
	"Darby Bible Translation":   DBT,
}
