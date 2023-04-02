package main

import "fmt"

func main() {
	fmt.Println("Break:")
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println("Index:", i)
	}

	fmt.Println("Continue:")
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}

		fmt.Println("Index:", i)
	}
}
