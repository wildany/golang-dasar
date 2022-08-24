package main

import "fmt"

func main() {
	counter := 0
	increament := func() {
		fmt.Println("Increament")
		counter++
	}

	increament()
	increament()
	fmt.Println(counter)
}
