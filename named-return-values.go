package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Mochammad"
	middleName = "Wildany"
	lastName = "Sihab"

	return firstName, middleName, lastName
}

func main() {

	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName, middleName, lastName)
}
