package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)


func main() {

	start := time.Now()

	cores := 2
	// if cores<0 then it will set to max cores no
	runtime.GOMAXPROCS(cores)

	var wg sync.WaitGroup

	for i:=0; i<1000000; i++{
		wg.Add(1)
		go print(i,&wg)
	}
	wg.Wait();

	end := time.Since(start)
	fmt.Println("Execution time :",end)
}

func print(n int, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println(n)
}