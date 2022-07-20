package main

import "fmt"

func getGoodBye(name string) string {
	return "Goof Bye " + name
}

func main() {
	sayGoodBye := getGoodBye

	result := sayGoodBye("Wildan")
	fmt.Println(result)
}
