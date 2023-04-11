package helper

import "fmt"

var version = 1                    //if variable or function start with lowercase then it cannot be accessed from outside
var Application = "Learn Golang" //if variable or function start with uppercase then it can be accessed from outside

func SayHello(name string) {
	fmt.Println("Hello", name)
}

func sayGoodbye(name string) {
	fmt.Println("Goodbye", name)
}
