package main

import "fmt"

func main() {
	// значение по умолчанию
	var num0 int

	// значение при инициализации
	var num1 int = 1

	// пропуск типа
	var num2 = 2

	// короткое объявление
	// только для новых переменных
	num3 := 3

	fmt.Println(num0, num1, num2, num3)

	num3 += 1

	num3++ // только пост инкремент

	// принятый стиль - CamelCase
	// объявление нескольких переменных
	var variableOne, variableTwo = 1, 2
	fmt.Println(variableOne, variableTwo)

	// короткое присваивание, хотя бы одна переменная должна быть новой
	variableOne, variableThree := 3, 4
	fmt.Println(variableOne, variableThree)

	variableOne, variableTwo = 3, 4
	fmt.Println(variableOne, variableTwo)

	var flag bool
	fmt.Println(flag)

	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64
	// float32 float64
	// bool
	// complex64 complex128
}
