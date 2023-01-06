package main

import "fmt"

// User
type User struct {
	nombre string
	email  string
	activo bool
}

// Student
type Student struct {
	user   User
	codigo string
}

func main() {
	alex := User{
		nombre: "Alex",
		email:  "alex@gmail.com",
		activo: true,
	}

	seba := User{
		nombre: "Seba",
		email:  "seba@gmail.com",
		activo: true,
	}

	alexStudent := Student{
		user:   alex,
		codigo: "001ARS",
	}

	fmt.Println(seba)
	fmt.Println(alexStudent.user.nombre)

}
