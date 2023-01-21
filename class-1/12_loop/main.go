package main

import "fmt"

func main() {

	// бесконечный цикл
	for {
		fmt.Println("endless")
		break
	}

	// цикл по переменной
	var stop = false
	for !stop {
		fmt.Println("endless")
		stop = true
	}

	// цикл с условием и инициализацией
	for i := 0; i < 2; i++ {
		fmt.Println("foo", i)
		if i == 1 {
			continue
		}
	}

Loop:
	// цикл с условием и инициализацией
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if i == 1 {
				break Loop
			}
			fmt.Println(i, j)
		}
	}

	slice := []int{1, 2, 3}
	idx := 0
	for idx < len(slice) {
		fmt.Println(slice[idx])
		idx++
	}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for idx := range slice {
		fmt.Println(slice[idx])
	}

	for idx, val := range slice {
		fmt.Println(idx, val)
	}

	for _, val := range slice {
		fmt.Println(val)
	}

	// операции по map
	foo := map[int]int{1: 2, 3: 4}
	for key := range foo {
		fmt.Println(key)
		foo[key+1] = key
	}

	for key, val := range foo {
		fmt.Println(key, val)
	}
	for _, val := range foo {
		fmt.Println(val)
	}

	// итерация по строке
	str := "Привет, Мир!"
	for pos, char := range str {
		fmt.Println(pos, string(char))
	}
}
