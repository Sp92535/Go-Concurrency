package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	var wg sync.WaitGroup

	wg.Add(2)

	go getNum(ch1, &wg)
	go getNum(ch2, &wg)

	go func() {
		wg.Wait()
		close(ch1)
		close(ch2)
	}()

	for ch1 != nil || ch2 != nil {
		select {
		case msg1, ok := <-ch1:
			if ok {
				fmt.Println("Channel 1 :", msg1)
			} else {
				ch1 = nil
			}
		case msg2, ok := <-ch2:
			if ok {
				fmt.Println("Channel 2 :", msg2)
			} else {
				ch2 = nil
			}
		}

	}

}

func getNum(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		ch <- rand.Intn(9)
	}
}
