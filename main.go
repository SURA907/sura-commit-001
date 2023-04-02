package main

import (
	"fmt"
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
		// rand.Seed(time.Now().UnixNano())
		// randM := rand.Intn(101)
		// t = t.Add(time.Duration(randM) * time.Second)

		fmt.Println("d:", t.Format(time.RFC3339))
	}

	return nil
}
