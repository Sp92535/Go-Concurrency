package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {

	ch := make(chan int, 3)
	var wg sync.WaitGroup

	wg.Add(1)
	go getNum(ch, &wg)
	wg.Wait()

	for i := range ch {
		fmt.Println(i)
	}

}

func getNum(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- rand.Intn(8)
	ch <- rand.Intn(8)
	ch <- rand.Intn(8)
	// ch <- rand.Intn(8)	// This will throw a deadlock due to the fixed size

	close(ch)
}
