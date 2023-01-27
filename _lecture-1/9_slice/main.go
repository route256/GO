package main

import "fmt"

func main() {
	// создание
	var a1 []int             // len=0, cap=0
	a2 := []int{}            // len=0, cap=0
	a3 := []int{1}           // len=1, cap=1
	a4 := make([]int, 0)     // len=0, cap=0
	a5 := make([]int, 5)     // len=5, cap=5
	a6 := make([]int, 5, 10) // len=5, cap=10

	fmt.Println(a1, a2, a3, a4, a5, a6)

	// обращение к элементам
	fmt.Println(a3[0])

	// ошибка в рантайме - panic: runtime error: index out of range [1] with length 1
	// fmt.Println(a3[1])

	var a []int          // len=0,cap=0
	a = append(a, 9, 10) // len=2,cap=2
	a = append(a, 16)    // len=3,cap=4

	// добавление другого слайса
	b := []int{1, 2, 3}
	a = append(a, b...)

	fmt.Println(
		len(a),
		cap(a),
	)

	sl1 := a[1:3] // [10,16]
	sl2 := a[:3]  // [9,10,16]
	sl3 := a[1:]  // [10,16,1,2,3]
	fmt.Println(sl1, sl2, sl3)

	newA := a[:] // [9,10,16,1,2,3]
	newA[0] = 8  // a = [8,10,16,1,2,3]
	fmt.Println(newA, a)

	// newA теперь указывает на другие данные
	newA = append(newA, 6, 7, 8)
	newA[0] = 6 // a = [8,10,16,1,2,3]
	fmt.Println(newA, a)

	// неправильное копирование
	var emptyA1 []int // len=0, cap=0
	copied := copy(emptyA1, a)
	fmt.Println(copied, emptyA1)

	// правильное копирование
	var emptyA2 = make([]int, len(a)) // len=0, cap=0
	copied = copy(emptyA2, a)
	fmt.Println(copied, emptyA2)

	// частичное копирование
	ints := []int{1, 2, 3, 4}
	copy(ints[1:3], []int{5, 6}) // ints = [1,5,6,4]
	fmt.Println(ints)

	// nil
	var nilSlice []int = nil
	fmt.Println(len(nilSlice))
	nilSlice = append(nilSlice, 5)
	fmt.Println(len(nilSlice))
}
