package main

import (
	"fmt"
	"sync"
)

var common int
var mu sync.Mutex

// go run -race main.go
func main() {
	var wg sync.WaitGroup

	wg.Add(10)

	for i := 0; i < 10; i++ {
		go inc(&wg)
	}

	wg.Wait()

	fmt.Println(common)
}

func inc(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	common++
	mu.Unlock()
}
