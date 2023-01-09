package main

import "fmt"

// Struct Persona
type Persona struct {
	nombre string
	edad   int
}

// Métodos Struct Persona
func (p *Persona) imprimir() {
	fmt.Printf("Nombre: %s\t Edad: %d\n", p.nombre, p.edad)
}

func (p *Persona) editNombre(nuevoNombre string) {
	p.nombre = nuevoNombre
}

// Herencia
type Empleado struct {
	Persona
	sueldo float32
}

func main() {
	p1 := Persona{"Seba", 35}
	p1.imprimir()
	p1.editNombre("Roberto")
	p1.imprimir()

	p2 := Persona{
		nombre: "Alex",
		edad:   27,
	}
	p2.imprimir()

	empl1 := Empleado{sueldo: 500}
	empl1.editNombre("Pedro")
	empl1.edad = 30
	empl1.imprimir()
	fmt.Println(empl1)
}
