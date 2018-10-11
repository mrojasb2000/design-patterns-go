package main

import (
	"fmt"
)

func main() {
	mySlice := make([]int, 10) // Create an underlying array of the ten elements.
	fmt.Println(len(mySlice), ":", cap(mySlice))
	mySlice = append(mySlice, 5) // Change the size of the Slice, add a new number.
	fmt.Println(len(mySlice), ":", cap(mySlice))

	mySlice = mySlice[1:] // Delete first element from the Slice

	mySlice = append(mySlice[:1], mySlice[2:]...)
}
