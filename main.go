package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"github.com/shelovesmox/minx-aio/checker"
)

func main() {
	// Number of requests to send
	numRequests := 2000

	// WaitGroup to synchronize goroutines
	var wg sync.WaitGroup
	wg.Add(numRequests)

	// Set GOMAXPROCS to the number of CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Start the timer
	startTime := time.Now()

	// Send multiple requests concurrently
	for i := 0; i < numRequests; i++ {
		go func() {
			defer wg.Done()

			// Send GET request
			resp, err := http.Get("https://example.com")
			if err != nil {
				log.Println("Request failed:", err)
				return
			}
			defer resp.Body.Close()
			// Read and discard the response body
			_, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println("Failed to read response body:", err)
			}

			// Increment the counter atomically
			atomic.AddUint64(&checker.Cpm, 1)

			fmt.Println("Request completed")
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Calculate the elapsed time
	elapsedTime := time.Since(startTime)

	// Get the final count
	count := atomic.LoadUint64(&checker.Cpm)

	fmt.Println("All requests completed")
	fmt.Println("Total requests:", count)
	fmt.Println("Elapsed time:", elapsedTime)
}
