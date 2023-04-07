package main

import "fmt"

func test(value int) interface{} {
	if value == 1 {
		return 1
	} else if value == 2 {
		return true
	} else {
		return "Not 1 or 2"
	}
}

func main() {
	var tryTest interface{} = test(2)
	fmt.Println(tryTest)
}
