package main

import (
	"fmt"
	"github.com/mrojasb2000/design-patterns-go/chapter1/aritmetic"
)

func main() {
	sumRes := aritmetic.Sum(5, 6)
	subRes := aritmetic.Subtract(10, 5)
	multiplyRes := aritmetic.Multiply(8, 7)
	divideRes, _ := aritmetic.Divide(10, 2)

	fmt.Printf("5+6 is %d. 10-5 is %d, 8*7 is %d, 10/2 is %f\n", sumRes, subRes, multiplyRes, divideRes)
}
