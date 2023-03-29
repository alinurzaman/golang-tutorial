package main

import "fmt"

func main() {
	name := "Muhamad"

	switch name {
	case "Ali":
		fmt.Println("Hello Ali")
	case "Nur Zaman":
		fmt.Println("Hello Nur Zaman")
	default:
		fmt.Println("Hello")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Data length is too long.")
	case length > 5:
		fmt.Println("Data length is ok.")
	default:
		fmt.Println("Data length is too short.")
	}
}
