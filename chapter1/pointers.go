package main

func main() {
	number := 5
	pointer_to_number := &number
	println(pointer_to_number)  // Address pointer
	println(*pointer_to_number) // Value on address
}
