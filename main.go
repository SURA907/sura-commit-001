package main

import (
	"errors"
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
	rand.Seed(time.Now().UnixNano())
	// 10 ~ 20
	commitCount := rand.Intn(10) + 10

	for i := 0; i < commitCount; i++ {
		_, err := os.Stat(filePath)

		fmt.Println(err)

		if err != nil && !errors.Is(err, os.ErrNotExist) {
			fmt.Println("err:", err)
			return err
		}

		if err == nil {
			// remove file
			fmt.Println("remove:", filePath)
			if errRemote := os.Remove(filePath); errRemote != nil {
				return errRemote
			}
		}

		if err == os.ErrNotExist {
			// create file
			fmt.Println("create:", filePath)
			if _, errCreate := os.Create(filePath); errCreate != nil {
				return errCreate
			}
		}

		cmdAdd := exec.Command("git", "add", ".")
		if errAdd := cmdAdd.Run(); errAdd != nil {
			return fmt.Errorf("git add: %v", errAdd)
		}

		commitMessage := fmt.Sprintf("commit -- %d", i)

		commitDate := d.Add(time.Duration(i) * time.Minute).Format(time.RFC3339)
		cmdCommit := exec.Command("git", "commit", "-m", commitMessage, "--date", commitDate)

		if errCommit := cmdCommit.Run(); errCommit != nil {
			return fmt.Errorf("git commit: %v", errCommit)
		}
	}

	os.Remove(filePath)

	return nil
}
