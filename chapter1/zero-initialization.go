package main

import (
	"errors"
	"log"
)

func main() {
	res, err := divisibleBy(10, 0)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v\n", res)
}

func divisibleBy(n, divisor int) (bool, error) {
	if divisor == 0 {
		return false, errors.New("A number cannot be divided by zero")
	}
	return (n%divisor == 0), nil
}
