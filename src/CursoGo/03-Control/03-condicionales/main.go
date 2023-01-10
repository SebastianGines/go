package main

import (
	"fmt"
)

func main() {
	/** App para restaurante
	* Descuentos de 10% hasta 100 de consumo
	* Descuentos de 20% hasta 200 de consumo
	* Descuentos de 30% mayor 200 de consumo
	* igv 19%
	 */
	var consumo, descuento float64
	var datosDescuento string

	// Entrada de datos
	fmt.Print("Ingrese consumo: ")
	fmt.Scanln(&consumo)

	if consumo >= 0 {
		if consumo <= 100 {
			descuento = 0.10
		} else if consumo <= 200 {
			descuento = 0.20
		} else {
			descuento = 0.30
		}
		// Operaciones
		datosDescuento = fmt.Sprintf("%.2f", descuento*100) + "%"
		descuento = descuento * consumo
		montoDescuento := consumo - descuento
		igv := montoDescuento * 0.19
		totalPagar := montoDescuento + igv
		// Salida de datos
		fmt.Println("================= Factura de Consumo =================")
		fmt.Println("\tDescuento que se aplica:", datosDescuento)
		fmt.Println("\tConsumo:", consumo)
		fmt.Println("\tDescuento:", descuento)
		fmt.Println("\tMonto con descuento:", montoDescuento)
		fmt.Println("\tIGV:", igv)
		fmt.Println("======================================================")
		fmt.Println("\tTotal a pagar:", totalPagar)
	} else {
		fmt.Println("Ingrese un consumo positivo")
	}
	fmt.Println()
}
