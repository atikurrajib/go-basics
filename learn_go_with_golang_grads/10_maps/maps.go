package main 

import (
	"fmt"
)

// maps -> hash, object, dict
func main() {

	// creating a map
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
}