package main

import "fmt"

func main() {
	name := "joko"

	if name == "dany" {
		fmt.Println("Hallo " + name)
	} else if name == "joko" {
		fmt.Println("hallo joko")
	} else {
		fmt.Println("Tanpa nama")
	}
}
