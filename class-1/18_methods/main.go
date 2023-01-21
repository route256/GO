package main

import "fmt"

type Person struct {
	ID   int
	Name string
}

type Account struct {
	ID   int
	Name string
	Person
}

// не изменяет оригинальную структуру
func (p Person) UpdateName(name string) {
	p.Name = name
}

// изменяет оригинальную структуру
func (p *Person) SetName(name string) {
	p.Name = name
}

// изменяет оригинальную структуру
func (a *Account) SetName(name string) {
	a.Name = name
}

func main() {
	pers := Person{1, "max"}

	pers.SetName("MAX")
	fmt.Println(pers)

	acc := Account{ID: 5, Name: "max"}
	acc.SetName("MAX")
	fmt.Println(acc)

	acc.Person.SetName("max")
	fmt.Println(acc)
	return
}
