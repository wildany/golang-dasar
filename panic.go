package main

import "fmt"

//panic adalah funtion yang digunakan untuk menghentikan program

func endApp() {
	fmt.Println("Aplikasi selesai")
}

func runtApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}

	fmt.Println("Aplikasi berjalan")
}

func main() {
	runtApp(true)
}
