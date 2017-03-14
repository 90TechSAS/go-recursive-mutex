package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/90TechSAS/go-recursive-mutex"
)

func main() {
	const Count = 10000000                     // Nb loops for testing
	const Workers = 1000                       // Nb Workers
	var timer time.Time                        // Timer
	var nativeMutex sync.Mutex                 // Native mutex
	var recursiveMutex recmutex.RecursiveMutex // Recursive mutex
	var workers = make(chan bool, Workers)     // Workers
	var counter int64                          // Counter used for testing concurrency
	var wg sync.WaitGroup                      // Wait group for syncing goroutines

	// Benchmarking native mutex
	wg.Add(Count)
	timer = time.Now()
	for i := 0; i < Count; i++ {
		workers <- true
		go func() {
			nativeMutex.Lock()
			counter++
			nativeMutex.Unlock()
			<-workers
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("[Native Mutex] Elapsed: %s / %d locks/sec / counter: %d\n", time.Since(timer), uint64(Count/time.Since(timer).Seconds()), counter)

	// Benchmarking recursive mutex
	counter = 0
	wg.Add(Count)
	timer = time.Now()
	for i := 0; i < Count; i++ {
		workers <- true
		go func() {
			recursiveMutex.Lock()
			counter++
			recursiveMutex.Unlock()
			<-workers
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("[Recursive Mutex] Elapsed: %s / %d locks/sec / counter: %d\n", time.Since(timer), uint64(Count/time.Since(timer).Seconds()), counter)
}
