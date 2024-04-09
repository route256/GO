package main

import "fmt"

// обычная функция
func nop() {
	fmt.Println("nop")
}

func main() {

	// анонимная функция
	func(in string) {
		fmt.Println(in)
	}("nop")

	// кладем функцию в переменную
	f := func(in string) {
		fmt.Println(in)
	}

	f("foo")

	// функция как тип
	type printer func(string)

	// функция принимает колбек
	worker := func(callback printer) {
		callback("from worker")
	}

	worker(f)

	// замыкание
	prefixer := func(prefix string) printer {
		return func(in string) {
			fmt.Printf("%s - %s\n", prefix, in)
		}
	}

	logger := prefixer("дата лога")
	logger("ошибка")
	logger("все ок")
}
