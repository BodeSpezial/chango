package files

import (
	"io"
	"os"
    "fmt"
    "time"
	"bufio"
    "strings"
	"io/ioutil"
)

//TODO give file as pointer instead of path
func file2lines(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return linesFromReader(f)
}

func linesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// Extracts and returns a section suurounded by the given parameters
func extractSubSection(section,begin string){
//var subsection string
fmt.Println(strings.Split(section,begin))
}

func between(value, a, b string) string {
fmt.Println(value)
fmt.Println(a)
fmt.Println(b)
    // Get substring between two strings.
    posFirst := strings.Index(value, a)
    if posFirst == -1 {
        return ""
    }

    posLast := strings.Index(value, b)
    if posLast == -1 {
        return ""
    }
    
    posFirstAdjusted := posFirst + len(a)
    if posFirstAdjusted >= posLast {
        return ""
    }
    fmt.Println(value[posFirstAdjusted:posLast])
    return value[posFirstAdjusted:posLast]
}

//TODO give file as pointer instead of path
func InsertStringToFile(path, str, targetLine string) error {
	lines, err := file2lines(path)
	if err != nil {
		return err
	}

today:=time.Now()
between(strings.Join(lines, "\n"), fmt.Sprintf("%v-%v-%v", today.Year(), int(today.Month()), today.Day()), "## ")

	fileContent := ""
	for _, line := range lines {
		if line == targetLine {
        //if strings.Contains(line, fmt.Sprintf("%v-%v-%v", today.Year(), int(today.Month()), today.Day())){
			fileContent += str
			continue
		}
		fileContent += line
		fileContent += "\n"
	}

	return ioutil.WriteFile(path, []byte(fileContent), 0644)
}
