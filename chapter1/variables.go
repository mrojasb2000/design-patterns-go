package main

import (
	"fmt"
	"reflect"
)

func main() {
	var explicit string = "Hello, I'm a explicitly declared"

	inferred := ", I'm an inferred variable "

	fmt.Println("Variable 'explicit' is of type:", reflect.TypeOf(explicit))
	fmt.Println("Variable 'inferred' is of type:", reflect.TypeOf(inferred))
}
