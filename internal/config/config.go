package config

import (
	"chango/structs"
	"github.com/BurntSushi/toml"
	"io/ioutil"

	"log"
)

var Config = structs.Configuration{}

func ParseConfigFile() *structs.Configuration {
	file, err := ioutil.ReadFile("./chango.toml")
	if err != nil {
		log.Fatal(err)
	}

	if _, err := toml.Decode(string(file), &Config); err != nil {
		log.Fatal(err)
	}

	return &Config
}
