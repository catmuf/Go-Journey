package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const numWorkers = 2              // Number of Go routines, matching the number of cores
	const maxCount = 1_000_000_000_00 // 1 trillion
	countPerWorker := maxCount / numWorkers

	var wg sync.WaitGroup
	results := make(chan int64, numWorkers)

	// Start the stopwatch
	start := time.Now()

	// Start counting in multiple Go routines
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int64, start, end int64) {
			defer wg.Done()
			var count int64
			for j := start; j < end; j++ {
				count++
			}
			results <- count
		}(int64(i), int64(i)*int64(countPerWorker), int64(i+1)*int64(countPerWorker))
	}

	// Wait for all Go routines to complete
	go func() {
		wg.Wait()
		close(results)
	}()

	// Sum up all the counts
	var totalCount int64
	for result := range results {
		totalCount += result
	}

	// Stop the stopwatch
	duration := time.Since(start)

	fmt.Printf("Total count: %d\n", totalCount)
	fmt.Printf("Time taken: %v\n", duration)
}
