package main

import "fmt"

/**
struct is like a class in oop
*/

type User struct {
	Name, Address string
	Age           int
}

func main() {
	var user User
	user.Name = "Muhamad Ali"
	user.Address = "Surabaya"
	user.Age = 29

	fmt.Println(user.Name)
	fmt.Println(user.Address)
	fmt.Println(user.Age)

	userTwo := User{
		Name:    "Bro",
		Address: "Malang",
		Age:     20,
	}
	fmt.Println(userTwo)

	userThree := User{"Sis", "Bandung", 25}
	fmt.Println(userThree)
}
