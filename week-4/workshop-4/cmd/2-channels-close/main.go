package main

import "fmt"

func main() {
	numbers := make(chan int, 16)

	go func() {
		defer close(numbers)

		for i := 0; i < 100; i++ {
			numbers <- i + 1
		}
	}()

	for i := range numbers {
		fmt.Println(i)
	}
}
