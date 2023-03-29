package main

import "fmt"

func main() {
	// map must add different key for each element
	person := map[string]string{
		"name":    "Ali",
		"address": "Surabaya",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Go-Lang"
	book["author"] = "Ali"
	book["price"] = "0"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "price")

	fmt.Println(book)
	fmt.Println(len(book))
}
