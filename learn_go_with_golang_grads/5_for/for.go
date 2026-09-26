package main

import (
	"fmt"
)

func main() {

	// for is only construct in go for looping

	/* 	// while with for loop:
	   	i := 1
	   	for i <= 3 {
	   		fmt.Println(i)
	   		i = i + 1
	   	} */

	/* 	infinite loop
	for {
		println("1")
	} */

	/* 	// classic for loo
	   	for i := 0; i <= 3; i++ {
	   		fmt.Println(i)
	   	} */

	/* 	// loop with break:
	   	for j := 1; j <= 9; j++ {
	   		break
	   		fmt.Println(j)
	   	} */

	/* 	// loop with continue
	   	for k := 0; k <= 9; k++ {
	   		if k == 6 {
	   			continue
	   		}
	   		fmt.Println(k) // print 0 to 9 with skiping 6
	   	} */

		// range
		for i := range 3 {
			fmt.Println(i)
		}

}
