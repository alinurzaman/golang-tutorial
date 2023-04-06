package main

import "fmt"

func logging() {
	fmt.Println("Done")
}

func runApp(value int) {
	defer logging()
	fmt.Println("Run application")
	result := 10 / value
	fmt.Println("Result", result)
}

func main() {
	runApp(0)
}
