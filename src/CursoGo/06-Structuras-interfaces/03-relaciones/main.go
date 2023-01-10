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

// Relación 1 a N (muchos)
type Curso struct {
	titulo string
	videos []Video
}

type Video struct {
	titulo string
	curso  Curso
}

func main() {
	/* Relación 1 a 1
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
	*/

	// Relación de 1 a muchos
	video1 := Video{titulo: "01-Introducción"}
	video2 := Video{titulo: "02-Instalación"}

	curso := Curso{
		titulo: "Curso de Go",
		videos: []Video{video1, video2},
	}

	video1.curso = curso
	video2.curso = curso

	fmt.Println(video1.curso.titulo)

	for _, video := range curso.videos {
		fmt.Println(video.titulo)
	}

}
