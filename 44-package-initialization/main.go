package main

import (
	"fmt"
	"golang-tutorial/database"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
