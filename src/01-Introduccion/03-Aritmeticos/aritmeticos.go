package main

import "fmt"

func main() {
	a := 10
	b := 20

	// Suma
	result := a + b
	fmt.Println("Suma: ", result)

	// Resta
	result = a - b
	fmt.Println("Resta: ", result)

	// Multiplicación
	result = a * b
	fmt.Println("Multiplicación: ", result)

	// División 1
	result = b / a
	fmt.Println("División: ", result)

	// División 2
	var div float32 = 10.0 / 20.0
	fmt.Println("División: ", div)

	// Módulo
	result = a % b
	fmt.Println("Módulo: ", result)
}
