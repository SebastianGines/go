package main

import (
	"fmt"
	"paquetes/models"

	figuras "github.com/SebastianGines/figurasGo"
	"github.com/donvito/hellomod"
)

func main() {
	// Módulo desde git
	cuadrado := figuras.Cuadrado{Alto: 2, Ancho: 2}
	circulo := figuras.Circulo{Radio: 3}
	fmt.Println("============================")
	figuras.CalcularArea(&cuadrado)
	figuras.CalcularPerimetro(&cuadrado)
	fmt.Println("============================")
	figuras.CalcularArea(&circulo)
	figuras.CalcularPerimetro(&circulo)

	// Módulo desde paquetes
	persona1 := models.Persona{}
	persona1.Constructor("Seba", 35)
	fmt.Println(persona1)
	persona1.SetNombre("Sebastian")
	fmt.Println(persona1)

	// Módulo desde git
	hellomod.SayHello()
}
