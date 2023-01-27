package main

func main() {

}

// обычное объявление
func singleIn(in int) int {
	return in
}

// много параметров
func multiIn(a, b, c int) int {
	return a + b + c
}

// именованная результат
func namedReturn() (out int) {
	out = 2
	return
}

// несколько результатов
func multipleReturn() (res int, err error) {
	return
}

// нефиксированное количество параметров
func sum(in ...int) (res int) {
	for _, val := range in {
		res += val
	}

	return
}
