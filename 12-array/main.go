package main

import "fmt"

func main() {
	var names [4]string

	names[0] = "Muhamad"
	names[1] = "Ali"
	names[2] = "Nur"
	names[3] = "Zaman"

	fmt.Println(names[0])
	fmt.Println(names[2])

	var values = [3]int{
		90, 95, 80,
	}

	fmt.Println(values)
	fmt.Println(values[1])

	fmt.Println(len(names))
	fmt.Println(len(values))

	var test [10]string
	fmt.Println(len(test))
}
