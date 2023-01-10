package main

import "fmt"

func main() {
	hola := "Hola"
	mundo := "Mundo"

	fmt.Println(hola)
	fmt.Println(mundo)

	nombre := "Seba"
	edad := 35

	fmt.Printf("Hola %s, su edad es %d \n", nombre, edad)
	fmt.Printf("Hola %v, su edad es %v \n", nombre, edad) // %v es cuando no sabemos que tipo de dato es

	mensaje := fmt.Sprintf("Hola %s, su edad es %d \n", nombre, edad)

	fmt.Println(mensaje)
	fmt.Printf("tipo de nombre: %T \n", nombre)

	fmt.Print("Ingrese otro nombre: ")
	fmt.Scanln(&nombre)

	fmt.Println("Otro nombre: ", nombre)

}
