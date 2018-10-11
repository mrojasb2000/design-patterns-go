package main

import (
	"fmt"
)

func main() {
	my_array := []int{1, 2, 3}
	for i := 0; i <= 10; i++ {
		println(i)
	}

	for index, value := range my_array {
		fmt.Printf("Index is %d and value is %d\n", index, value)
	}

	for index := 0; index < len(my_array); index++ {
		value := my_array[index]
		fmt.Printf("Index is %d and value is %d\n", index, value)
	}
}
