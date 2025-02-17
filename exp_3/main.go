	package main

	import (
		"fmt"
		"math/rand"
		"sync"
	)

	func main() {

		ch := make(chan int)
		var wg sync.WaitGroup

		var z int

		fmt.Print("Enter total random numbers desired :")
		fmt.Scan(&z)

		wg.Add(z)

		for i:=0; i<z; i++{
			go getNum(ch,&wg)
		}
		
		go func(){
			wg.Wait()
			close(ch)
		}()
		
		for val := range ch{
			fmt.Println(val)
		}
		
		
	}

	func getNum(ch chan<- int, wg *sync.WaitGroup) {
		ch <- rand.Intn(9)
		defer wg.Done()
	}
