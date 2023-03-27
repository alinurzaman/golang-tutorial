package main

import "fmt"

func main() {
	var firstName = "Muhamad"
	var lastName = "Ali"

	var result bool = firstName == lastName
	fmt.Println(result) // false

	var i = 100
	var j = 500
	fmt.Println(i > j)  // false
	fmt.Println(i < j)  // true
	fmt.Println(i == j) // false
	fmt.Println(i != j) // true

}
