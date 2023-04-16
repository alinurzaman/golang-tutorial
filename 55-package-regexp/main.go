package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("a([a-z])i")

	fmt.Println(regex.MatchString("ali"))
	fmt.Println(regex.MatchString("api"))
	fmt.Println(regex.MatchString("aDi"))

	search := regex.FindAllString("ali alo aka api ani are", -1)
	fmt.Println(search)
}
