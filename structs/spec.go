package structs

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

type Configuration struct {
	Editor      string
	RepoUrl     string
	ChlogFolder string `toml:"path"`
	AutoCommit  bool   `toml:"auto_commit"`
}
