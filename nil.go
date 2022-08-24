package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}
func main() {
	var person map[string]string = newMap("Dany")

	if person == nil {
		fmt.Print("Data kosong")
	} else {
		fmt.Print(person)
	}

}
