package main

import "fmt"

func main() {
	// длина массива является частью типа и задается при компиляции
	var a1 [3]int // [0,0,0]
	fmt.Println("a1 short", a1)
	fmt.Printf("a1 short %v\n", a1)
	fmt.Printf("a1 short %#v\n", a1)

	// можно использовать константы для указания размера
	const size = 2
	var a2 [2 * size]bool // [false,false,false,false]
	fmt.Println("a2", a2)

	// определение размера при объявлении
	a3 := [...]int{1, 2, 3}
	fmt.Println("a3", a3)

	// проверка при компиляции - invalid argument: index 4 out of bounds [0:3]
	// a3[4] = 12
	// проверка в рантайме - panic: runtime error: index out of range [5] with length 3
	idx := 5
	a3[idx] = 12

}
