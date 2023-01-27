package main

import "fmt"

type ID int

func main() {
	var id ID = 5
	fmt.Println(id)

	foo(int(id))
}

func foo(bar int) {
	fmt.Println(bar)
}
