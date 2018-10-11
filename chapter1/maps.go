package main

import (
	"fmt"
)

func main() {
	myMap := make(map[string]int)
	myMap["one"] = 1
	myMap["two"] = 2

	fmt.Println(myMap["one"])
}
