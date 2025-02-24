package main

import (
	"fmt"
	"sync"
)

var shared int = 9
var mutex sync.RWMutex

func main() {
	var wg sync.WaitGroup
	
	wg.Add(3)

	go read(&wg)
	go write(&wg)
	go read(&wg)
	
	defer wg.Wait()
}

func read(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.RLock()
	fmt.Println(shared)
	mutex.RUnlock()
}
func write(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	shared++
	mutex.Unlock()
}
