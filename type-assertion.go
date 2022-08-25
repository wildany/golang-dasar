package main

import (
	"fmt"
)

func random() interface{} {
	return "Ups"
}

func main() {
	var result interface{} = random()
	// var resulString string = result.(string)
	// fmt.Print(resulString)

	switch value := result.(type) {
	case string:
		fmt.Print("Value ", value, " is string")
	case int:
		fmt.Print("Value ", value, " is int")
	default:
		fmt.Print("Unknown Type")
	}

}
