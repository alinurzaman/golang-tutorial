package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello Guest!"
	} else {
		return "Hello " + name + "!"
	}
}
func main() {
	result := getHello("Ali")
	fmt.Println(result)

	fmt.Println(getHello(""))
}
