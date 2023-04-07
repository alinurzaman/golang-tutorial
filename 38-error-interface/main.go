package main

import (
	"errors"
	"fmt"
)

func calculate(value int, divider int) (int, error) {
	if divider == 0 {
		return 0, errors.New("divider shouldn't be 0")
	} else {
		result := value / divider
		return result, nil
	}
}
func main() {
	result, e := calculate(100, 0)
	if e == nil {
		fmt.Println(result)
	} else {
		fmt.Print(e.Error())
	}
}
