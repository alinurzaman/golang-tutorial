package main

import "fmt"

/**
it's like a method of class in oop
*/

type User struct {
	Name, Address string
	Age           int
}

func (user User) sayHello(name string) {
	fmt.Println("Hello", name, "My Name is", user.Name)
}

func (user User) sayHelloTwo() {
	fmt.Println("Hello from:", user.Name, user.Address, user.Age)
}

func main() {
	var user User
	user.Name = "Muhamad Ali"
	user.Address = "Surabaya"
	user.Age = 29

	user.sayHello("Guest")
	user.sayHelloTwo()
}
