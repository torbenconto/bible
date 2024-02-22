package main

type version struct {
	name string
	path string
	url  string
}

var versions = []version{
	version{"kjv", "~/.bible/versions/kjv/", "https://openbible.com/textfiles/kjv.txt"},
}
