# Bible CLI
A simple command line interface for the Bible.

## Installation
```bash
 go install github.com/torbenconto/bible
```

## Usage
Flags:
`--version <Bible Version (eg: KJV or ASV)>` pulls from that version of the bible

### Get verse by name
```bash
bible get "John 3:16"
```

### Get multiple verses
```bash
bible get "John 3:16-17"
```

### Compare verses in different versions
```bash
bible compare "John 3:16-7" ASV KJV (etc)
```
Takes N versions and will print out the verse(s) in each version

### Get Random Verse
```bash
bible random
```

### Get Random Verse from a specific book
```bash
bible random "John"
```