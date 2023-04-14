package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Muhamad Ali", "Ali"))
	fmt.Println(strings.Contains("Muhamad Ali", "Nur"))

	fmt.Println(strings.Split("Muhamad Ali Nur Zaman", " "))

	fmt.Println(strings.ToLower("Muhamad Ali Nur Zaman"))
	fmt.Println(strings.ToUpper("Muhamad Ali Nur Zaman"))
	fmt.Println(strings.ToTitle("muhamad ali nur zaman"))

	fmt.Println(strings.Trim("      Muhamad Ali Nur     ", " "))
	fmt.Println(strings.ReplaceAll("Ali Nur Ali", "Ali", "Muhamad"))
}
