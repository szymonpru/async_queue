package tasks

import (
	"log"
	"time"
)

const sleepSeconds = 2

func Sum(numbers []int64) (int64, error) {
	log.Printf("Starting task 'Add'. Numbers to sum: %v\n", numbers)

	var sum int64 = 0
	for _, arg := range numbers {
		sum += arg
		time.Sleep(sleepSeconds * time.Second)
	}

	log.Println("Task finished.")
	return sum, nil
}
