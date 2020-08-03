package files

import (
	"io"
	"io/ioutil"
    "bufio"
    "os"
)

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

func InsertStringToFile(path, str, targetLine string) error {
	lines, err := file2lines(path)
	if err != nil {
		return err
	}

	fileContent := ""
	for _, line := range lines {
		if line == targetLine {
			fileContent += str
			continue
		}
		fileContent += line
		fileContent += "\n"
	}

	return ioutil.WriteFile(path, []byte(fileContent), 0644)
}
