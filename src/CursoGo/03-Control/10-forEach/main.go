package main

import "fmt"

func main() {
	nombres := [...]string{"Alex", "Roel", "Juan"}

	// for normal
	for i := 0; i < len(nombres); i++ {
		fmt.Println(nombres[i])
	}

	// for each mostrando también índice
	for indice, elemento := range nombres {
		fmt.Println(indice, elemento)
	}

	// for each sin mostrar el índice
	for _, elemento := range nombres {
		fmt.Println(elemento)
	}

	// for each sin mostrar el elemento
	for indice := range nombres {
		fmt.Println(indice)
	}
}
