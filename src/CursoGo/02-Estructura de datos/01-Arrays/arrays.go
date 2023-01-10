package main

import "fmt"

func main() {
	// Arrays
	var numeros [5]int
	fmt.Println(numeros)
	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	numeros[3] = 40
	numeros[4] = 50
	fmt.Println(numeros)

	// Arrays con valores
	nombres := [3]string{"Alex", "Roel", "Juan"}
	fmt.Println(nombres)
	colores := [...]string{
		"Rojo",
		"Verde",
		"Negro",
		"Azul",
	}
	fmt.Println(colores, len(colores))

	// Indices definidos
	monedas := [...]string{
		0: "Dolar",
		2: "Soles",
		3: "Pesos",
		5: "Euro",
	}
	fmt.Println(monedas, len(monedas))

	// SubArrays
	subMonenda := monedas[:4]
	fmt.Println(subMonenda, len(subMonenda))
	subMonenda2 := monedas[2:]
	fmt.Println(subMonenda2, len(subMonenda2))
}
