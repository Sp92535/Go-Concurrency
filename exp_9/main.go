package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 8)
	var wg sync.WaitGroup

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel() // Ensures cleanup

	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go printNum(ctx, i, &wg, ch)
	}

	// Separate goroutine to wait and exit properly on timeout
	go func() {
		wg.Wait()
		cancel() // Extra safety to stop context if not already done
	}()

	<-ctx.Done() // Main function exits after timeout
	fmt.Println("Timeout reached. Exiting.")
}

func printNum(ctx context.Context, n int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	select {
	case <-ctx.Done(): // Check if timeout occurred
		return
	case ch <- 1: // Proceed if not timed out
		fmt.Println(n)
		<-ch
	}
}
