package main

import "fmt"

func sumar(nombre string, numeros ...int) (string, int) {
	mensaje := fmt.Sprintf("La suma de %s es:", nombre)
	suma := 0
	for _, numero := range numeros {
		suma += numero
	}
	return mensaje, suma
}

func main() {
	mensaje, result := sumar("Seba", 1, 2, 5, 4, 0)
	fmt.Println(mensaje, result)
}
