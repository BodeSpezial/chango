package markdown

import (
	"chango/structs"

	"fmt"
	"sort"
	"time"
	"strconv"
)

var prNumbers []int

const repoUrl = "https://github.com/BodeSpezial/chango/pulls"

// addSection adds a section with a list of changes in that section
func addSection(name string, s *[]structs.ChangelogEntry, m *string) {
	if len(*s) > 0 {
		*m += "### " + name + "\n"
		for _, section := range *s {
			*m += "  - " + section.Message + " (#" + strconv.Itoa(section.Pr) + ")\n"
		}
		*m += "\n"
	}
}

// concatPrLink creates the link to the pr that is needed in the 'Version Links'-section
func concatPrLink(number int, repoUrl string) string {
	return "[#" + strconv.Itoa(number) + "]: " + repoUrl + "/" + strconv.Itoa(number)
}

// listVersionLinks groups all the pr numbers to a list of links
func appendVersionLinks(m *string) error {
	for _, num := range prNumbers {
		versionLink := "[#" + strconv.Itoa(num) + "]: " + repoUrl + "/" + strconv.Itoa(num) + "\n"
		*m += versionLink
	}
	return nil
}

// appendPrNumbers takes a section of the changelog and appends it to the list of pr numbers
func appendPrNumbers(s *[]structs.ChangelogEntry) error {
	for _, section := range *s {
		prNumbers = append(prNumbers, section.Pr)
	}
	return nil
}

// generateMarkdown generates the final markdown to write to the file
func GenerateMarkdown(changes structs.Changelog) string {
	md := ""
	date := time.Now()

	md += fmt.Sprintf("## %v-%v-%v\n\n", date.Year(), int(date.Month()), date.Day())

	addSection("Added", &changes.Added, &md)
	addSection("Changed", &changes.Changed, &md)
	addSection("Fixed", &changes.Fixed, &md)
	addSection("Removed", &changes.Removed, &md)

	md += "\n<!--Version Links-->\n\n"

	appendPrNumbers(&changes.Added)
	appendPrNumbers(&changes.Changed)
	appendPrNumbers(&changes.Fixed)
	appendPrNumbers(&changes.Removed)

	sort.Ints(prNumbers)
	appendVersionLinks(&md)

	return md
}
