package git

import (
	"log"
	"os/exec"
)

func GetCurrentBranch() string {
	branch, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		log.Println(err)
	}
	branchName := string(branch)
	return branchName
}

