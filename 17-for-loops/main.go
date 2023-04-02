package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Counter:", counter)
		counter++
	}

	for i := 1; i <= 10; i++ {
		fmt.Println("i:", i)
	}

	numbers := []string{"One", "Two", "Three"}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	for index, number := range numbers {
		fmt.Println(index, ":", number)
	}

	// if you dont need to access the index you can use '_'
	for _, number := range numbers {
		fmt.Println(number)
	}

	person := make(map[string]string)
	person["name"] = "Ali"
	person["title"] = "Engineer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
