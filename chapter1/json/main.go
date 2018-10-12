package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	packt := "packt"
	jsonPackt, ok := json.Marshal(packt)
	if ok != nil {
		panic("Could not marshal object")
	}
	fmt.Println(string(jsonPackt))
}
