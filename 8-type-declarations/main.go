package main

import "fmt"

func main() {
	type idNumber string

	var ktpAli idNumber = "123456789"
	fmt.Println(ktpAli)
	fmt.Println(idNumber("987654321"))
}
