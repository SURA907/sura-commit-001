package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const DATE_FORMAT = "2006-01-02 14:05:06"

func main() {
	if err := run(); err != nil {
		fmt.Println("ERR:", err)
	}
}

func run() error {
	fmt.Println("----> start")

	t := time.Date(2023, time.January, 01, 12, 11, 0, 0, time.Local)

	for !time.Now().Before(t) {
		t = t.AddDate(0, 0, 1)
		rand.Seed(time.Now().UnixNano())
		randM := rand.Intn(101)

		d := t.Add(time.Duration(randM) * time.Minute)
		if err := createCommits(d); err != nil {
			return err
		}
	}

	return nil
}

func createCommits(d time.Time) error {
	filePath := fmt.Sprintf("create-%d", time.Now().UnixMilli())

	fmt.Println(filePath)
	rand.Seed(time.Now().UnixNano())
	// 10 ~ 20
	commitCount := rand.Intn(10) + 10

	for i := 0; i < commitCount; i++ {
		_, err := os.Stat(filePath)
		if err != nil && err != os.ErrNotExist {
			return err
		}

		if err == nil {
			// remove file
			if errRemote := os.Remove(filePath); errRemote != nil {
				return errRemote
			}
		}

		if err == os.ErrNotExist {
			// create file
			if _, errCreate := os.Create(filePath); errCreate != nil {
				return errCreate
			}
		}

		cmdAdd := exec.Command("git", "add", ".")
		if errAdd := cmdAdd.Run(); errAdd != nil {
			return errAdd
		}

		commitMessage := fmt.Sprintf("commit -- %d", i)
		

		cmdCommit := exec.Command("git", "commit", "-m", commitMessage, "--date=")

	}

	return nil
}
