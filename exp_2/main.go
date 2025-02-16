package main

import (
	"fmt"
	"sync"
)


func main() {	
	var wg sync.WaitGroup
	
	go printNumbers(4, 11, &wg)
	go printNumbers(8, 15, &wg)
	wg.Add(2)
	
	wg.Wait()
}

func printNumbers(l int, r int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := l; i <= r; i++ {
		fmt.Println(i)
	}
}
