package main

import (
	"fmt"
)

func main() {
	numbers := make(chan int)

	go func() {
		fmt.Println("Waiting for a number in numbers channel")
		v := <-numbers
		fmt.Println("Read number: ", v)
	}()

	numbers <- 1
	numbers <- 2
}
