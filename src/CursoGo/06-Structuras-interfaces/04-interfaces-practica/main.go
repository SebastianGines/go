package main

import (
	"fmt"
	"math"
	"reflect"
)

type Figura interface {
	area() float32
	perimetro() float32
}

type Cuadrado struct {
	Alto  float32
	Ancho float32
}

type Circulo struct {
	radio float32
}

func (cuadrado *Cuadrado) area() float32 {
	return cuadrado.Alto * cuadrado.Ancho
}

func (cuadrado *Cuadrado) perimetro() float32 {
	return 2*cuadrado.Alto + 2*cuadrado.Ancho
}

func (circulo *Circulo) area() float32 {
	return math.Pi * (circulo.radio * circulo.radio)
}

func (circulo *Circulo) perimetro() float32 {
	return 2 * math.Pi * circulo.radio
}

func calcularArea(figura Figura) {
	fmt.Println("El área de:", reflect.TypeOf(figura).Elem().Name(), "es:", figura.area())
}

func calcularPerimetro(figura Figura) {
	fmt.Println("El perímetro de", reflect.TypeOf(figura).Elem().Name(), "es:", figura.perimetro())
}

func main() {
	cuadrado := Cuadrado{2, 2}
	circulo := Circulo{3}
	fmt.Println("============================")
	calcularArea(&cuadrado)
	calcularPerimetro(&cuadrado)
	fmt.Println("============================")
	calcularArea(&circulo)
	calcularPerimetro(&circulo)
}
