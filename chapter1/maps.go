package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	myMap := make(map[string]int)
	myMap["one"] = 1
	myMap["two"] = 2

	fmt.Println(myMap["one"])

	myJsonMap := make(map[string]interface{})
	jsonData := []byte(`{"hello":"world"}`)
	err := json.Unmarshal(jsonData, &myJsonMap)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", myJsonMap["hello"])
}
