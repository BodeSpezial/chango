package changes

import (
	"chango/structs"
)
// gatherChanges reads all the files from the changelog folder and returns them in as a Changelog
func GatherChanges(path string) structs.Changelog {
	if path == "" {
		path = "./changelog/"
	}
	changes := structs.Changelog{}

	return changes
}
