package main

import "fmt"

func checkType() interface{} {
	return "Ali" // try and change this return
}

func main() {
	var result interface{} = checkType()

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown type")
	}
}
