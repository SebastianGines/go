package main

import "fmt"

func main() {
	// Not
	fmt.Println(!true)
	// And
	fmt.Println(false && true)
	// Or
	fmt.Println(true || false)

	b := 2
	r := b == 2 || !(4 > b)
	println(r)
}
