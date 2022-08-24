package main

import "fmt"

func main() {
	counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke ", counter)
	// 	counter++
	// }

	for i := 0; i < 20; i++ {
		fmt.Println("Perulangan ke ", counter)
	}

	slice := []string{"mochammad", "wildany", "sihab"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("index", i, "=", value)
	}

	for _, value := range slice {
		fmt.Println(value)
	} // _ digunakan untuk memberitahu golang jika _ tidak digunakan

	person := make(map[string]string)
	person["name"] = "Dany"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}

}
