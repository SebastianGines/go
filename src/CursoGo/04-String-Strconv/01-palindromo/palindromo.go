package main

import (
	"fmt"
	"strings"
)

func reverse(cadena string) string {
	arrayCadena := strings.Split(cadena, "")
	arrayInvertida := make([]string, 0)
	for i := len(arrayCadena) - 1; i >= 0; i-- {
		arrayInvertida = append(arrayInvertida, arrayCadena[i])
	}
	return strings.Join(arrayInvertida, "")
}

func esPalindromo(palabra string) bool {
	palabra = strings.ToLower(palabra)
	palabra = strings.ReplaceAll(palabra, " ", "")
	palabraInvertida := reverse(palabra)
	return palabra == palabraInvertida
}

func main() {
	fmt.Printf("Ingrese una palabra para evaluar si es palíndromo: ")
	var palabra string
	fmt.Scanln(&palabra)
	if esPalindromo(palabra) {
		fmt.Println("Es palíndromo")
	} else {
		fmt.Println("No es palíndromo")
	}
}
