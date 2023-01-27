package main

import "fmt"

func testPanic() {
	defer func() {
		if err := recover(); err != nil {
			panic(err)
		}
	}()

	fmt.Println("work")
	panic("oops")
	fmt.Println("work again")
	return
}

func main() {
	testPanic()
}
