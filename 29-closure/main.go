package main

import "fmt"

func main() {
	name := "Ali"
	counter := 0

	increment := func() {
		name := "Zaman"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
