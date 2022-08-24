package main

//recover adalah function yang digunakan untuk menangkap nilai panik. jika ada fungsi recover, maka aplikasi tidak jadi berhenti

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runtApp(err bool) {
	defer endApp()
	if err {
		panic("APLIKASI ERROR")
	}

	fmt.Println("Aplikasi berjalan")
}

func main() {
	runtApp(true)
	println("dany")
}
