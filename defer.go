package main

import "fmt"

//defer = fungsi yang dipanggil walaupun terjadi error dan selalu di eksekusi terakhir
func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("Run application")
	result := 10 / value
	fmt.Println("Result", result)
}
func main() {
	runApplication(10)
}
