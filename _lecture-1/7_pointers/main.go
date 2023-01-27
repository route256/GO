package main

import "fmt"

func main() {
	a := 2
	b := &a
	*b = 3  // a = 3
	c := &a // новый указатель на переменную

	// получение указателя на переменную типа int
	// инициализировано значением по умолчанию
	d := new(int)
	*d = 12
	*c = *d // c = 12 -> a = 12
	*d = 13

	c = d   // теперь c = 13
	*c = 14 // c = 14 -> d = 14, a = 12

	var e *int
	fmt.Println(e)

	c = nil
	fmt.Println(c)
	return
}
