package main

import (
	"fmt"
)

// slices -> dynamic
// most used construct in go
// usefull methods

func main() {

/* 	// uninitialized slice is nil
	var nums []int
	fmt.Println(nums)
	fmt.Println(len(nums))
	fmt.Println(nums == nil) */

 	var nums = make([]int , 2)
	fmt.Println(nums)
	fmt.Println(cap(nums))
	fmt.Println(len(nums))

	nums = append(nums, 5)
	nums = append(nums, 4)
	nums = append(nums, 3)

	fmt.Println(nums)
	fmt.Println(cap(nums))
	fmt.Println(len(nums))

/*  var names =  make([]string, 5)
	fmt.Println(names)
	fmt.Println(cap(names))
	fmt.Println(len(names))

	names = append(names, "baron")
	names = append(names, "carolina")
	names = append(names, "hobert")

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

}
