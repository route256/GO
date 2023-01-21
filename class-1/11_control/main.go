package main

import "fmt"

func main() {
	// условие
	if true {
		fmt.Println("true")
	}

	mapVal := map[string]string{"name": "max"}
	// условие с блоком иницализации
	if val, ok := mapVal["name"]; ok {
		fmt.Println(val)
	}

	// только проверяем наличие ключа
	if _, ok := mapVal["name"]; ok {
		fmt.Println("ключ существует")
	}

	cond := 1
	// if else
	if false {
		fmt.Println("false")
	} else if 1 == 2 {
		fmt.Println("foo")
	} else if cond == 1 {
		fmt.Println(cond)
	} else {
		fmt.Println("else")
	}

	// switch по 1 переменной
	str := "name"
	switch str {
	case "name":
		fallthrough
	case "test", "lastName":
		fmt.Println("test")
	default:
		// noop
	}

	// switch как замена ifelse
	var val1, val2 = 2, 2
	switch {
	case val1 > 1 || val2 < 11:
		fmt.Println("one")
	case val2 > 10:
		fmt.Println("two")
	}

	return
}
