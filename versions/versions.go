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

var AMP = *NewVersion("amp", "amp.txt", "https://raw.githubusercontent.com/torbenconto/bibles/master/amp.txt", []string{})
var NIV = *NewVersion("niv", "niv.txt", "https://raw.githubusercontent.com/torbenconto/bibles/master/niv.txt", []string{})
var ESV = *NewVersion("esv", "esv.txt", "https://raw.githubusercontent.com/torbenconto/bibles/master/esv.txt", []string{})
var KJV = *NewVersion("kjv", "kjv.txt", "https://openbible.com/textfiles/kjv.txt", []string{"KJV", "King James Bible: Pure Cambridge Edition - Text courtesy of www.BibleProtector.com"})
var NKJV = *NewVersion("nkjv", "nkjv.txt", "https://raw.githubusercontent.com/torbenconto/bibles/master/nkjv.txt", []string{})
var ASV = *NewVersion("asv", "asv.txt", "https://openbible.com/textfiles/asv.txt", []string{"ASV", "American Standard Version"})
var BSB = *NewVersion("bsb", "bsb.txt", "https://bereanbible.com/bsb.txt", []string{"The Holy Bible, Berean Standard Bible, BSB is produced in cooperation with Bible Hub, Discovery Bible, OpenBible.com, and the Berean Bible Translation Committee.", "This text of God's Word has been dedicated to the public domain. Free resources and databases are available at BereanBible.com.\t", "Verse\tBerean Standard Bible"})
var NASB = *NewVersion("nasb", "nasb.txt", "https://raw.githubusercontent.com/torbenconto/bibles/master/nasb2020.txt", []string{})
var DBT = *NewVersion("dbt", "dbt.txt", "https://openbible.com/textfiles/dbt.txt", []string{"DBT", "Darby Bible Translation"})
var DRB = *NewVersion("drb", "drb.txt", "https://openbible.com/textfiles/drb.txt", []string{"DRB", "Douay-Rheims Bible"})
var ERV = *NewVersion("erv", "erv.txt", "https://openbible.com/textfiles/erv.txt", []string{"ERV", "English Revised Version"})
var YLT = *NewVersion("ylt", "ylt.txt", "https://openbible.com/textfiles/ylt.txt", []string{"YLT", "Young's Literal Translation"})
var NLT = *NewVersion("nlt", "nlt.txt", "https://raw.githubusercontent.com/torbenconto/bibles/master/nlt.txt", []string{})
var EASY = *NewVersion("easy", "easy.txt", "https://raw.githubusercontent.com/torbenconto/bibles/master/easy.txt", []string{})
var CSB = *NewVersion("csb", "csb.txt", "https://raw.githubusercontent.com/torbenconto/bibles/master/csb.txt", []string{})

var Versions = []Version{
	AMP,
	NIV,
	ESV,
	KJV,
	NKJV,
	ASV,
	BSB,
	NASB,
	DBT,
	DRB,
	ERV,
	YLT,
	NLT,
	EASY,
	CSB,
}

var RecommendedVersions = []Version{
	KJV,
	NKJV,
	NIV,
	ESV,
}

var VersionMap = map[string]Version{
	"AMP":                         AMP,
	"Amplified Bible":             AMP,
	"NIV":                         NIV,
	"New International Version":   NIV,
	"ESV":                         ESV,
	"English Standard Version":    ESV,
	"KJV":                         KJV,
	"King James Version":          KJV,
	"NKJV":                        NKJV,
	"New King James Version":      NKJV,
	"ASV":                         ASV,
	"American Standard Version":   ASV,
	"BSB":                         BSB,
	"Berean Standard Bible":       BSB,
	"NASB":                        NASB,
	"New American Standard Bible": NASB,
	"DBT":                         DBT,
	"Darby Bible Translation":     DBT,
	"DRB":                         DRB,
	"Douay-Rheims Bible":          DRB,
	"ERV":                         ERV,
	"English Revised Version":     ERV,
	"YLT":                         YLT,
	"Young's Literal Translation": YLT,
	"NLT":                         NLT,
	"New Living Translation":      NLT,
	"EASY":                        EASY,
	"Easy-to-Read Version":        EASY,
	"CSB":                         CSB,
	"Christian Standard Bible":    CSB,
}
