package main

func main() {
	addN := func(m int) func(int) int {
		return func(n int) int {
			return m + n
		}
	}

	//addFive := addN(5)
	//result := addFive(6)
	result := addN(5)(6)
	// 5 + 6 must print 11
	println(result)

}
