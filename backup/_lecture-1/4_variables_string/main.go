package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// пустая строка по умолчанию
	var one string

	// со спец символами
	var two = "Hello 世界\n"

	// с экранированием
	var three = `
		Hello World\n
	`

	fmt.Println(one, two, three)

	// одинарные кавычки для байт (uint8)
	var rawByte byte = '\x27'
	fmt.Println((rawByte))

	// rune (uint32) для UTF-8 символов
	var rawRune rune = 'c'
	fmt.Println((rawRune))

	// конкантенация строк
	fmt.Println(two + three)

	// строки неизменяемы
	// three[0] = 72

	// длина строки в байтах
	fmt.Println(len(two))
	// длина строки в символах
	fmt.Println(utf8.RuneCountInString(two))

	// срез строки в байтах, не символах
	fmt.Println(two[:9])

	// в слайс байт и обратно
	byteString := []byte(two)
	two = string(byteString)
	fmt.Println(two)
	return
}
