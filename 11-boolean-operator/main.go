package main

import "fmt"

func main() {
	var condition1 = 80
	var condition2 = 75

	var passCondition1 = condition1 >= 80
	var passCondition2 = condition2 >= 80
	fmt.Println(passCondition1) // true
	fmt.Println(passCondition2) // false

	fmt.Println(condition1 >= 80 || condition2 >= 80) // true
	fmt.Println(condition1 >= 80 && condition2 >= 80) // false
}
