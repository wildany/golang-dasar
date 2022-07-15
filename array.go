package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Mochammad"
	names[1] = "Wildany"
	names[2] = "Sihab"

	fmt.Println(names[0] + names[1] + names[2])

	var values = [3]int{
		1,
		2,
		3,
	}

	fmt.Println(values)
}
