package main

import "fmt"

func main() {
	// инициализация
	var user = map[string]string{
		"name":    "Maxim",
		"surname": "Pak",
	}
	fmt.Println(user["name"])

	// nil
	// var user1 map[string]string
	// fmt.Println(len(user1))
	// user1["name"] = "Maxim"
	// fmt.Println(user1["name"])

	// сразу с нужным размером
	var profile = make(map[string]string, 10)
	fmt.Printf("%d %+v\n", len(profile), profile)

	// если ключа нет, вернет значение по умолчанию
	mName := user["middleName"]
	fmt.Println("mName:", mName)

	// проверка на существование ключа
	mName, ok := user["middleName"]
	fmt.Println("mName:", mName, "ok:", ok)

	// пустая переменная
	_, ok = user["middleName"]
	fmt.Println("ok:", ok)

	// удаление ключа
	delete(user, "surname")
	fmt.Printf("%#v\n", user)

	delete(user, "surname")
	fmt.Printf("%#v\n", user)

}
