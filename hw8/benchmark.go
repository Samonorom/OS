// benchmark.go
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var levels = []string{"INFO", "WARN", "ERROR"}
var contexts = []string{"req-123", "user-456", "req-789", "user-999"}

func generateRandomLogEntry() LogEntry {
	return LogEntry{
		Timestamp: time.Now(),
		Level:     levels[rand.Intn(len(levels))],
		Context:   contexts[rand.Intn(len(contexts))],
		Message:   fmt.Sprintf("Message %d", rand.Intn(1000)),
	}
}

func runBenchmark(name string, logger interface {
	Log(LogEntry) error
	Close()
}) time.Duration {
	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 50; j++ {
				entry := generateRandomLogEntry()
				logger.Log(entry)
			}
		}(i)
	}

	wg.Wait()
	logger.Close()
	return time.Since(start)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Running benchmarks for 5 goroutines × 50 log entries each...")

	// Naive Logger
	if logger, err := NewNaiveLogger("naive.log"); err == nil {
		duration := runBenchmark("NaiveLogger", logger)
		fmt.Printf("NaiveLogger took: %v\n", duration)
	} else {
		fmt.Println("NaiveLogger error:", err)
	}

	// Mutex Logger
	if logger, err := NewMutexLogger("mutex.log"); err == nil {
		duration := runBenchmark("MutexLogger", logger)
		fmt.Printf("MutexLogger took: %v\n", duration)
	} else {
		fmt.Println("MutexLogger error:", err)
	}

	// Channel Logger
	if logger, err := NewChannelLogger("channel.log"); err == nil {
		duration := runBenchmark("ChannelLogger", logger)
		fmt.Printf("ChannelLogger took: %v\n", duration)
	} else {
		fmt.Println("ChannelLogger error:", err)
	}
}
