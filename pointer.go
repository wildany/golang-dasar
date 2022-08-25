package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}
func main() {
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// address2 := address1
	address2 := &address1
	address4 := new(Address)
	address2.City = "Bandung"
	// address1.City = "Surabaya"
	// address2 = &Address{"Surabaya", "Jawa Timur", "Indonesia"}
	*address2 = Address{"Surabaya", "Jawa Timur", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address4)

	var alamat = Address{
		City:     "Malang",
		Province: "Jawa Timur",
		Country:  "",
	}
	ChangeCountryToIndonesia(&alamat)
	fmt.Print(alamat)
}
