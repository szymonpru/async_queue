package tasks

import (
	"fmt"
	"time"
)

const sleepSeconds = 10

func DummyTask() error {
	fmt.Printf("Starting task. Sleeping for %v seconds. \n", sleepSeconds)
	time.Sleep(sleepSeconds * time.Second)
	fmt.Println("Task finished.")
	return nil
}
