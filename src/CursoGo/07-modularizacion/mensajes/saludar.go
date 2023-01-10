package mensajes

import "fmt"

func Hola() {
	fmt.Println("Hola desde el paquete mensaje")
}

const mensaje = "Hola desde mi constante"

func Imprimir() {
	fmt.Println(mensaje)
	funcionPrivada()
}

func funcionPrivada() {
	fmt.Println("Hola desde la función privada")
}
