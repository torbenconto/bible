package main

type Version struct {
	Name     string
	Path     string
	Url      string
	BadLines []string
}

func NewVersion(name, path, url string, badLines []string) *Version {
	return &Version{name, path, url, badLines}
}

var KJV = *NewVersion("kjv", "kjv.txt", "https://raw.githubusercontent.com/scrollmapper/bible_databases/master/csv/t_kjv.csv", []string{"KJV", "King James Bible: Pure Cambridge Edition - Text courtesy of www.BibleProtector.com"})

var versions = []Version{
	KJV,
}
