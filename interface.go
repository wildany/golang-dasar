package main

import "fmt"

type HasName interface {
	getName() string
}

func sayHi(hasName HasName) {
	fmt.Println("Hello", hasName.getName())
}

type Person struct {
	Name string
}

func (person Person) getName() string {
	return person.Name
}

//contoh 2
type Animal struct {
	Name string
}

func (animal Animal) getName() string {
	return animal.Name
}
func main() {
	var dany Person
	dany.Name = "Dany"

	sayHi(dany)

	cat := Animal{
		Name: "Pussy",
	}
	sayHi(cat)
}
