package main

import "fmt"

func getMyName(name string) string {
	return "My name is " + name
}

func main() {
	sayMyName := getMyName
	result := sayMyName("Ali")
	fmt.Println(result)
	fmt.Println(getMyName("Ali"))
}
