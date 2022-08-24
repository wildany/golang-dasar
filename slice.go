package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))

	months[5] = "diubah"
	fmt.Println(slice1)

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "dany")
	fmt.Println(slice3)
	slice3[1] = "Bukan_Desember"
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)

	var tesCap = months[2:4]
	fmt.Println(cap(tesCap))

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Dany"
	newSlice[1] = "Sihab"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	iniArrat := [...]int{1, 2, 3, 4, 5} //ini Array, ... digunakan jika belum mengetahui kapasitasnya
	iniSlice := []int{1, 2, 3, 4, 5}

}
