package main

import "fmt"

func getName() (firstName, lastName string) {
	firstName = "Muhamad"
	lastName = "Ali"

	return
}

func main() {
	firstName, lastName := getName()
	fmt.Println(firstName, lastName)
}
