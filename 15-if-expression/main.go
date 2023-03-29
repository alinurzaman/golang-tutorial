package main

import "fmt"

func main() {
	name := "Ali"

	if name == "Ali" {
		fmt.Println("Hello")
	} else {
		fmt.Println("Hola")
	}

	// short statement condition
	if length := len(name); length > 5 {
		fmt.Println("Data failed")
	} else {
		fmt.Println("Data ok")
	}
}
