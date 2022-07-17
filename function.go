package main

import "fmt"

func sayHello(name string) {
	fmt.Println("hai", name)
}

func luasSegitiga(alas float64, tinggi float64) float64 {

	return alas * tinggi * 0.5
}
func main() {
	name := "Dany"
	var alas float64 = 12
	var tinggi float64 = 20

	sayHello(name)
	fmt.Println(luasSegitiga(alas, tinggi))

}
