package changes

import (
	"chango/internal/yaml"
	"chango/structs"
	"io/ioutil"
	"log"
	"os"
)

// gatherChanges reads all the files from the changelog folder and returns them in as a Changelog
func GatherChanges(path string) structs.Changelog {
	if path == "" {
		path = "./changelog/"
	}
	changes := structs.Changelog{}

	files, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer files.Close()
	list, _ := files.Readdirnames(0)
	for _, name := range list {

		file, err := os.Open(path + name)
		if err != nil {
			log.Fatal(err)
		}

		data, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}

		changes.Added = append(changes.Added, yaml.Yamlify(data).Added...)
		changes.Changed = append(changes.Changed, yaml.Yamlify(data).Changed...)
		changes.Fixed = append(changes.Fixed, yaml.Yamlify(data).Fixed...)
		changes.Removed = append(changes.Removed, yaml.Yamlify(data).Removed...)
	}

	return changes
}
