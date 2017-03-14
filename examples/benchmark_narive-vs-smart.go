package main

import (
	"fmt"
	"github.com/90TechSAS/go-smart-mutex"
	"sync"
	"time"
)

func main() {
	const Count = 10000000                 // Nb loops for testing
	const Workers = 1000                   // Nb Workers
	var timer time.Time                    // Timer
	var mutex sync.Mutex                   // Native mutex
	var smartMutex smutex.SmartMutex       // Smart mutex
	var workers = make(chan bool, Workers) // Workers
	var counter int64                      // Counter used for testing concurrency
	var wg sync.WaitGroup                  // Wait group for syncing goroutines

	// Benchmarking native mutex
	wg.Add(Count)
	timer = time.Now()
	for i := 0; i < Count; i++ {
		workers <- true
		go func() {
			mutex.Lock()
			counter++
			mutex.Unlock()
			<-workers
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("[Native Mutex] Elapsed: %s / %d locks/sec / counter: %d\n", time.Since(timer), uint64(Count/time.Since(timer).Seconds()), counter)

	// Benchmarking smart mutex
	counter = 0
	wg.Add(Count)
	timer = time.Now()
	for i := 0; i < Count; i++ {
		workers <- true
		go func() {
			smartMutex.Lock()
			counter++
			smartMutex.Unlock()
			<-workers
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("[Smart Mutex] Elapsed: %s / %d locks/sec / counter: %d\n", time.Since(timer), uint64(Count/time.Since(timer).Seconds()), counter)
}
