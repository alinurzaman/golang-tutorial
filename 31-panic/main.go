package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error with message:", message)
	}
	fmt.Println("Success")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Application error")
	}
	fmt.Println("Application running")
}

func main() {
	runApp(true)
}
