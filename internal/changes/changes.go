package changes

import (
	"chango/structs"
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

	}

	return changes
}
