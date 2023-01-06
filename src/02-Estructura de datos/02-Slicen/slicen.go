package main

import "fmt"

func main() {
	// slicen
	numeros := []int{1, 2, 3}
	fmt.Println(numeros, len(numeros))

	// Agregar datos
	numeros = append(numeros, 4)
	fmt.Println(numeros, len(numeros))

	// SubSlicen
	subSlicen := numeros[:2]
	numeros[0] = 100
	fmt.Println(subSlicen)
	fmt.Println(numeros)

	// Punteros
	// Capacidad
	// Longitud

	meses := []string{"Enero", "Febrero", "Marzo"}
	fmt.Printf("Len: %v, Cap: %v, %p \n", len(meses), cap(meses), meses)

}
