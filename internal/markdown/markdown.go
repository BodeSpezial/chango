package markdown

import (
	"chango/structs"
	"fmt"
	"sort"
	"strconv"
	"time"
)

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

