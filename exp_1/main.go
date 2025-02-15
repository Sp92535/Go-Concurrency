package main

import (
	"fmt"
	"time"
)

func main() {

	go printNumbers(4, 11)
	go printNumbers(8, 15)

	time.Sleep(time.Second)
}

func printNumbers(l int, r int) {
	for i := l; i <= r; i++ {
		fmt.Println(i)
	}
}
