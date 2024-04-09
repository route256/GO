package main

import "fmt"

const pi = 3.14
const e = 2.67

// enum
const (
	zero int = iota // автоинкремент
	one             //
	_               // пустая переменная, пропуск iota
	two             //
	three
	four
)

var (
	a = 1
	b = 2
)

const (
	_         = iota             // пропускаем 0
	Kb uint64 = 1 << (10 * iota) // 1024
	Mb                           // 1048576
)

const (
	foo     = 1 // нетипизированная константа
	bar int = 2 // типизированная константа
)

func main() {
	var hut int32 = 1
	fmt.Println(hut + foo)
	fmt.Println(pi, zero, one, two)
}
