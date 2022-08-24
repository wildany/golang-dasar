package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

//struct function

func (customer Customer) sayHai(name string) {
	fmt.Println("Hai", name, "My name is", customer.Name)
}
func main() {
	//cara 1
	var person Customer
	person.Name = "Mochammad Wildany Sihab"
	person.Address = "Indonesia"
	person.Age = 23

	//cara 2
	person2 := Customer{
		Name:    "Person 2",
		Address: "Indonesia",
		Age:     30,
	}

	//cara 3
	person3 := Customer{"Person3", "Indonesia", 31}

	fmt.Println(person)
	fmt.Println(person2)
	fmt.Println(person3)

	person.sayHai("Dany")
}
