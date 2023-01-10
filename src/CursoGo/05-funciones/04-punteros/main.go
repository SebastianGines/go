package main

import "fmt"

func modificarValor(cadena *string) {
	fmt.Printf("%p\n", cadena)
	*cadena = "Hola desde la función"
}

func main() {
	cadena := "Hola mundo de Go"
	fmt.Printf("%p\n", &cadena)
	fmt.Println("Antes de la función: ", cadena)
	modificarValor(&cadena) // "cadena" copia de la cadena \ "&cadena" la referencia de la cadena
	fmt.Println("Después de la función: ", cadena)
}
