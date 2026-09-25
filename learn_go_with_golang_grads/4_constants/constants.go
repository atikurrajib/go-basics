package main

import (
	"fmt"
)

var name1 string = "adam"

func main() {
	
	const name2 = "travis"

	fmt.Println(name1)
	fmt.Println(name2)

	const (
		port = 5000
		host = "local host"
	)

	fmt.Println(port, host)
}
