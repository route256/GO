package main

import "fmt"

type Person struct {
	ID      int
	Name    string
	Address string
}

type Account struct {
	ID    int
	Name  string
	Owner Person
	F     func(string) string
	Person
}

func main() {
	// полное объявление
	var acc = Account{}
	fmt.Println(acc)

	// краткое объявление
	acc.Owner = Person{2, "Max", "Moscow"}
	fmt.Println(acc)

	acc.Address = "Ryazan"
	fmt.Println(acc.Address)

	acc.Person.ID = 4
	fmt.Println(acc)
}
