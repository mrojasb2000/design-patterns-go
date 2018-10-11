package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%d\n", sum(1, 2, 3))
	fmt.Printf("%d\n", sum(4, 5, 6, 7, 8))
}

func sum(args ...int) (result int) {
	for _, v := range args {
		result += v
	}
	return
}
