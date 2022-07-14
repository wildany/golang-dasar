package main

import "fmt"

func main() {
	const phi = 3.14

	const (
		panjang = 10
		lebar   = 20
		jari    = 5
	)

	luasPersegi := panjang * lebar
	luasLingkaran := phi * jari * jari
	fmt.Println(luasLingkaran)
	fmt.Println(luasPersegi)

}
