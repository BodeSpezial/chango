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

func CommitChangelog(changelogFile string) error {
	addChlog := exec.Command("git", "add", changelogFile)
	if err := addChlog.Run(); err != nil {
		return err
	}

	commitChlog := exec.Command("git", "commit", "-m", "Added changelog")
	if err := commitChlog.Run(); err != nil {
		return err
	}
	return nil
}
