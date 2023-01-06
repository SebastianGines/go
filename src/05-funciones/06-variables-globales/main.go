package main

import "fmt"

// variable global
var mensaje string

func funcion1() {
	mensaje = "Hola desde la función 1"
	fmt.Println(mensaje)
}

func funcion2() {
	mensaje = "Hola desde la función 2"
	fmt.Println(mensaje)
}

func main() {
	mensaje = "Hola desde la función principal!"
	fmt.Println(mensaje)

	defer funcion1()
	funcion2()

	fmt.Println(mensaje)
}
