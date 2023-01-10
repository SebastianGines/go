package main

import "fmt"

func main() {
	numeros := make([]int, 3, 3)
	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30

	numeros = append(numeros, 40)

	fmt.Println(numeros, len(numeros), cap(numeros))
}
