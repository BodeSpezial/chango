package config

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"

	"log"
)

var Config = Configuration{}

func (c Configuration) UpdateConfig() {
	file, err := ioutil.ReadFile("./chango.toml")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := toml.Decode(string(file), &Config); err != nil {
		log.Fatal(err)
	}
}

type Configuration struct {
	Editor  string
	RepoUrl string
	Git     struct {
		ChlogFolder string `toml:"path"`
		AutoCommit  bool   `toml:"auto_commit"`
		CommitMsg   string `toml:"commit_message"`
	}
}
