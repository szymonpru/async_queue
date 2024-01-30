package tasks

import (
	"fmt"
	"time"
)

const sleepSeconds = 1

func DummyTask() (string, error) {
	fmt.Printf("Starting task. Sleeping for %v seconds. \n", sleepSeconds)
	time.Sleep(sleepSeconds * time.Second)
	fmt.Println("Task finished.")
	return "Result 123", nil
}
