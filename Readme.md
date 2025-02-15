# Go Concurrency

This repository covers various experiments to understand concurrency, multithreading, and parallelism in Go.

## Experiments

### 1. Goroutines and Basic Concurrency
**Problem Statement:** Create a program where multiple goroutines print numbers from different ranges concurrently.

### 2. Syncing Goroutines with WaitGroups
**Problem Statement:** Modify the previous experiment to ensure all goroutines finish before the main function exits.

### 3. Communication Between Goroutines (Channels)
**Problem Statement:** Implement a producer-consumer model using Go channels.

### 4. Buffered vs Unbuffered Channels
**Problem Statement:** Demonstrate the difference between buffered and unbuffered channels by implementing a message queue.

### 5. Select Statement for Multiple Channels
**Problem Statement:** Implement a system where multiple goroutines send messages to different channels, and the `select` statement picks available messages.

### 6. Mutex and Race Conditions
**Problem Statement:** Create a shared counter that multiple goroutines increment, leading to a race condition. Then, fix it using `sync.Mutex`.

### 7. Read-Write Locks (sync.RWMutex)
**Problem Statement:** Implement a program where multiple goroutines read a shared variable while only one can write to it at a time.

### 8. Worker Pool Pattern
**Problem Statement:** Create a worker pool where a fixed number of goroutines process multiple tasks from a queue.

### 9. Context for Goroutine Cancellation
**Problem Statement:** Implement a service that runs indefinitely until a timeout is reached, using `context.WithTimeout()`.

### 10. Parallelism with GOMAXPROCS
**Problem Statement:** Run computationally intensive tasks with different values of `runtime.GOMAXPROCS()` and measure performance.

---

## How to Run Experiments
1. Clone this repository.
2. Navigate to the desired experiment folder.
3. Run the Go program using `go run main.go`.

Happy coding! 🚀

