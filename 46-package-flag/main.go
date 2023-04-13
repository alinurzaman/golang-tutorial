package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "Put your database host")
	var user *string = flag.String("user", "root", "Put your database user")
	var password *string = flag.String("password", "root", "Put your database password")
	var number *int = flag.Int("number", 10, "Put your number")

	// go run main.go -host={hostName} -user={userName} -password={yourPassword} -number={number}

	flag.Parse()

	fmt.Println("Host:", *host)
	fmt.Println("User:", *user)
	fmt.Println("Password: ", *password)
	fmt.Println("Number:", *number)
}
