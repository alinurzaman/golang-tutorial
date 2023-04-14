package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true") //parsing string to boolean
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	number, err := strconv.ParseInt("1000000", 10, 64) //parsing string to int
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	value := strconv.FormatInt(1000000, 10) //format int to string
	fmt.Println(value)

	valueInt, _ := strconv.Atoi("2000000")
	fmt.Println(valueInt)
}
