package main

type Changelog struct {
	Added   []ChangelogEntry `yaml:"Added"`
	Fixed   []ChangelogEntry `yaml:"Fixed"`
	Changed []ChangelogEntry `yaml:"Changed"`
	Removed []ChangelogEntry `yaml:"Removed"`
}

type ChangelogEntry struct {
	Message string
	Pr      int
}

func main() {

}
