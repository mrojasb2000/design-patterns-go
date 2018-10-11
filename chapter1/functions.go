package main

import (
	"fmt"
)

func hello(message string) error {
	fmt.Printf("Hello %s\n", message)
	return nil
}

func main() {
	hello("Mavro")
}
