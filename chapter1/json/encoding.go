package main

import (
	"encoding/json"
	"fmt"
)

type myObject struct {
	Number int    `json:"number"`
	Word   string `json:"string"`
}

func main() {

	jsonBytes := []byte(`{"number":5, "string":"Packt"}`)
	var dangerousObject map[string]interface{}
	err := json.Unmarshal(jsonBytes, &dangerousObject)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Number is %f, ", dangerousObject["number"])
	fmt.Printf("Word is %s, ", dangerousObject["string"])
	fmt.Printf("Error reference is %v, ", dangerousObject["nothing"])

	/* 	jsonBytes := []byte(`{"number":5, "string":"Packt"}`)
	var object MyObject
	err := json.Unmarshal(jsonBytes, &object)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Number is %d, Word is %s\n", object.Number, object.Word)*/

	/*	object := MyObject{5, "Packt"}
		oJson, _ := json.Marshal(object)
		fmt.Printf("%s\n", oJson) */
}
