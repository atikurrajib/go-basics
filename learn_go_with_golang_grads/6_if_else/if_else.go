package main

import (
	"fmt"
)

func main() {

	/* 	age := 6

	   	if age >= 18 {
	   		fmt.Println("person is an adult")
	   	} else if age >= 13 {
	   		fmt.Println("person is a teenager")
	   	} else {
	   		fmt.Println("person is a kid")
	   	} */

	/* 	var name string
	   	var age int

	   	fmt.Println("Enter your name: ")
	   	fmt.Scan(&name)

	   	fmt.Println("Enter your age: ")
	   	fmt.Scan(&age)

	   	if age >= 18 {
	   		fmt.Printf("hi %s, you are %d years old. so, you are an adult.", name, age)
	   	} else if age >= 13 {
	   		fmt.Printf("hi %s, you are %d years old. so, you are a teenager.", name, age)
	   	} else {
	   		fmt.Printf("hi %s, you are %d years old. so, you are a kid.", name, age)
	   	} */

	var role = "admin"
	var hasPermissions = true

	if role == "admin" || hasPermissions {
		fmt.Println("yes")
	}

	// go doesn't have ternary, you will have to use normal if else
}
