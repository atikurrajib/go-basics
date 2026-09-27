package main 

import (
	"fmt"
)

// maps -> hash, object, dict
func main() {

/* 	// creating a map
	m := make(map[string]string) // varName := make(map[keytype]valuetype)
	// setting an element
	m["name"] = "golang"
	// get an element
	fmt.Println(m["name"])

	n := make(map[string]string)
	n["area"] = "backend"
	fmt.Println(m["name"], n["area"])

	// if key does not exists in the map (when string is valuetype) then it returns blank
	fmt.Println(m["course"])

	// if key does not exists in the map (when int is valuetype) then it returns zero
	a := make(map[string]int)
	a["phone"] = 2343552
	fmt.Println(a["contact"])

	// if key does not exists in the map (when bool is valuetype) then it returns false
	b := make(map[string]bool)
	b["isCompleted"] = true
	fmt.Println(b["courseCompleted"])
	
	fmt.Println(len(b))

	b["isFailed"] = false
	fmt.Println(len(b))

	// delete function
	delete(b, "isFailed")
	fmt.Println(b)
	fmt.Println(len(b))

	// clear function
	clear(b)
	fmt.Println(b) 

	m:= map[string]int{"price": 40, "phones": 3}
	fmt.Println(m) */

	m:= map[string]int{"price": 40, "phones": 3}
	fmt.Println(m)

	k, ok := m["phones"]
	fmt.Println((k))
	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}

}