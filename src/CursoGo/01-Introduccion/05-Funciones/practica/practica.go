package main

import "fmt"

func suma(a, b int) int {
	return a + b
}

func cociente(a, b int) int {
	return a / b
}

func residuo(a, b int) int {
	return a % b
}

func igv(a int, igv float32) float32 {
	return float32(a) * 0.18
}

func main() {
	var num1 int
	var num2 int
	fmt.Print("Ingrese un número: ")
	fmt.Scanln(&num1)
	fmt.Print("Ingrese otro número: ")
	fmt.Scanln(&num2)
	fmt.Println("La suma de", num1, "y", num2, "es", suma(num1, num2))
	fmt.Println("El cociente de", num1, "y", num2, "es", cociente(num1, num2))
	fmt.Println("El residuo de", num1, "y", num2, "es", residuo(num1, num2))
	fmt.Print("Ingrese el precio de venta del producto: ")
	fmt.Scanln(&num1)
	igv := igv(num1, 0.18)
	precioFinal := float32(num1) + igv
	fmt.Println("El IGV es", igv, "y el precio final es", precioFinal)
}
