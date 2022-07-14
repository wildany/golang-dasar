package main

import "fmt"

func main() {
	type noKTP string

	var noKtpDany noKTP = "12131313131312"
	fmt.Println(noKtpDany)
}
