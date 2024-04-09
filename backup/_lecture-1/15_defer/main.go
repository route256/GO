package main

import "fmt"

func main() {
	defer fmt.Println("after return 1")
	defer fmt.Println("after return 2")
	fmt.Println("before return")
}
