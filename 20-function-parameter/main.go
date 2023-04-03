package main

import "fmt"

func sayHello(name string, email string) {
	fmt.Println("Name:", name)
	fmt.Println("Email:", email)
}

func main() {
	name := "Ali"
	sayHello(name, "ali@mail.com")
	sayHello("Muhamad", "ali@mail.com")
}
