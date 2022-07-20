package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello ", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

//bisa juga menggunakan filter
//type Filter func[string] string
func main() {
	sayHelloWithFilter("wildan", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)

}
