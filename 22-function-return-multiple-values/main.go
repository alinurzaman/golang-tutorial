package main

import "fmt"

func getFullName() (string, string) {
	return "Muhamad", "Ali"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)

	newFirstName, _ := getFullName()
	fmt.Println(newFirstName)
}
