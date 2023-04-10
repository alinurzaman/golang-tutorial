package main

import "fmt"

type Person struct {
	Name string
}

func (person *Person) Married() {
	person.Name = "Mr. " + person.Name
}

func main() {
	ali := Person{"Ali"}
	ali.Married()

	fmt.Println(ali.Name)
}
