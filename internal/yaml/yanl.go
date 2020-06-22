package yaml

import (
	"chango/structs"
	"gopkg.in/yaml.v2"
	"log"
)

// yamlify unmarshals content to a Changelog var and returns it
func Yamlify(content []byte) structs.Changelog {
	chlog := structs.Changelog{}

	err := yaml.Unmarshal(content, &chlog)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return chlog
}
