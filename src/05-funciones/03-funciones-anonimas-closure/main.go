package main

import (
	"fmt"
	"strings"
)

// Closure

func repeat(n int) func(cadena string) string {
	return func(cadena string) string {
		return strings.Repeat(cadena, n)
	}
}

func main() {

	/*
		func() {
			fmt.Println("hola desde una función anónima")
		}()
	*/
	myfunc := func(nombre string) string {
		return fmt.Sprintf("Hola %s", nombre)
	}
	fmt.Println(myfunc("Seba"))

	repeat3 := repeat(3)
	fmt.Println(repeat3("hola"))

}
