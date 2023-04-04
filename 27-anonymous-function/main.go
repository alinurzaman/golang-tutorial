package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blackList BlackList) {
	if blackList(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blackList := func(name string) bool {
		return name == "admin"
	}
	registerUser("Ali", blackList)
	registerUser("admin", blackList)

	// you can initialize anonymous function on parameter
	registerUser("Ali", func(name string) bool {
		return name == "not user"
	})
	registerUser("not user", func(name string) bool {
		return name == "not user"
	})
}
