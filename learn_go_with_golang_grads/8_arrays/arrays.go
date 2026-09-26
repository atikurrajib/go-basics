package main

import (
	"fmt"
)

// number sequence of specific length
func main() {
	var nums [4]int

	nums[0] = 1

	/* 	// array length
	   	fmt.Println(len(nums)) */

	fmt.Println(nums[0])
	fmt.Println(nums) // int -> assign zero values by default

	var vals [4]bool
	fmt.Println(vals) // bool -> assign false by default

	var name [3]string
	name[0] = "golang"
	fmt.Println(name) // string -> assign empty string by default

	// declare it in single line, 1d array
	numbers := [3] int{1, 2, 3}  // varName := [size] dataType{data1, data2, data(size)}
	fmt.Println(numbers)

	// 2d array
	cars := [2][2] string{{"mercedes", "benz"}, {"supra", "buggati"}}
	fmt.Println(cars)
}
