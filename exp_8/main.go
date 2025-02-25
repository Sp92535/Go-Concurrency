package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	ch := make(chan int, 8 /*X*/)

	// fmt.Println(runtime.NumCPU())   If confused keep X = runtime.NumCPU()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go print(i, &wg, ch)
	}
	
	wg.Wait()
}

func print(num int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	ch <- 1
	fmt.Println(num)
	<-ch
}
