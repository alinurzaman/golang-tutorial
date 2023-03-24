package main

import "fmt"

func main() {
	var name string

	name = "Muhamad"
	fmt.Println(name)

	name = "Ali"
	fmt.Println(name)

	var newName = "Nur Zaman"
	fmt.Println(newName)

	var age = 30
	fmt.Println(age)

	country := "Indonesia"
	fmt.Println(country)

	var (
		city  = "Surabaya"
		state = "East Java"
	)
	fmt.Println(city, ",", state)
}
