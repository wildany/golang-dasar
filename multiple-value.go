package main

import "fmt"

func getFullName() (string, string) {
	return "wildany", "sihab"
}

func main() {
	firstName, lastName := getFullName()

	fmt.Println(firstName)
	fmt.Println(lastName)
}
