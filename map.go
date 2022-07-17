package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Dany",
		"address": "Sidoarjo",
	}

	person["age"] = "23"
	fmt.Println(person)
	fmt.Println(person["name"])

	var book = make(map[string]string)
	book["title"] = "Belajar golang"
	book["author"] = "Dany"
	book["Ups"] = "Salah"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "Ups")

	fmt.Println(book)
	fmt.Print(len(book))
}
