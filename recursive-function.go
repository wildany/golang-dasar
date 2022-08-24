package main

import "fmt"

func faktorial(value int) int {
	if value <= 0 {
		return 1
	}
	return value * faktorial(value-1)
}

func tailFaktorial(total int, value int) int {
	if value <= 0 {
		return total
	}
	return tailFaktorial(total*value, value-1)
}

func main() {
	fmt.Println(faktorial(5))
	fmt.Println(tailFaktorial(5, 4))
}
