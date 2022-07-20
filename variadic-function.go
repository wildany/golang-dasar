package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := sumAll(10, 10, 10, 10, 10)
	fmt.Println(total)

	//jika memasukan data berupa slice
	slice := []int{10, 20, 30, 40, 50}
	totalSlice := sumAll(slice...)
	fmt.Println(totalSlice)
}
