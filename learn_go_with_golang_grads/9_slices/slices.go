package main

import (
	"fmt"
	"slices"
)

// slices -> dynamic
// most used construct in go
// useful methods

func main() {

	/* 	// uninitialized slice is nil
	   	var numbers []int
	   	fmt.Println(numbers)
	   	fmt.Println(len(numbers))
	   	fmt.Println(numbers == nil) */

	/*  var numbers = make([]int , 2)
	fmt.Println(numbers)
	fmt.Println(cap(numbers))
	fmt.Println(len(numbers))

	numbers = append(numbers, 5)
	numbers = append(numbers, 4)
	numbers = append(numbers, 3)

	fmt.Println(numbers)
	fmt.Println(cap(numbers))
	fmt.Println(len(numbers)) */

	/*  var names =  make([]string, 5)
	fmt.Println(names)
	fmt.Println(cap(names))
	fmt.Println(len(names))

	names = append(names, "baron")
	names = append(names, "carolina")
	names = append(names, "robert")

	fmt.Println(names)
	fmt.Println(cap(names))
	fmt.Println(len(names)) */

	/* 	var isPassed =  make([]bool, 2)
	   	fmt.Println(isPassed)
	   	fmt.Println(cap(isPassed))
	   	fmt.Println(len(isPassed))

	   	isPassed = append(isPassed, true)
	   	isPassed = append(isPassed, true)
	   	isPassed = append(isPassed, true)

	   	fmt.Println(isPassed)
	   	fmt.Println(cap(isPassed))
	   	fmt.Println(len(isPassed)) */

	/* 	var numbers = make([]int, 2, 10)
	   	fmt.Println(numbers)
	   	fmt.Println(cap(numbers))
	   	fmt.Println(len(numbers))

	   	numbers = append(numbers, 2)
	   	fmt.Println(numbers)
	   	fmt.Println(cap(numbers))
	   	fmt.Println(len(numbers))

	   	numbers[0], numbers[1] = 0, 1
	   	fmt.Println(numbers)
	   	fmt.Println(cap(numbers))
	   	fmt.Println(len(numbers))

	   	// copy function
	   	numbers := make([]int, len(numbers))
	   	copy(numbers, numbers)
	   	fmt.Println(numbers) */

	/* 	// slice operator
	   	var numbers = []int{1, 2, 3}
	   	fmt.Println(numbers[0:2])
	   	fmt.Println(numbers[:2])
	   	fmt.Println(numbers[2:]) */

	// slice comparison
	var num1 = []int{1, 2}
	var num2 = []int{3, 4}

	fmt.Println(slices.Equal(num1, num2))

	// 2d slices
	var numbers = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(numbers)
}
