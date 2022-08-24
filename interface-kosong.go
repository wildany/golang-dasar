package main

import "fmt"

//digunakan jika ingin membuat parameter dg tipe data apapun
func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}
func main() {
	var data interface{} = Ups(2)

	fmt.Print(data)
}
