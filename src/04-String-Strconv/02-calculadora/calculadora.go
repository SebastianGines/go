package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sumar(expresion string) int {
	valores := strings.Split(expresion, "+")
	var result int

	for _, valor := range valores {
		num, error := strconv.Atoi(valor)

		if error != nil {
			fmt.Println(error)
			fmt.Println("Error: tiene que ingresar un número entero")
			fmt.Println("o solo debes realizar operaciones con el operador +")
		} else {
			result += num
		}
	}
	return result
}

func main() {
	var expresion string
	var result int

	fmt.Print("=> ")
	fmt.Scanln(&expresion)
	result = sumar(expresion)
	fmt.Printf("=> %d\n", result)
}
