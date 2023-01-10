package main

import (
	"fmt"
	"net/http"
	"time"
)

func revisarServidor(servidor string, canal chan string) {
	_, error := http.Get(servidor)

	if error != nil {
		//fmt.Println(servidor, "no está disponible")
		canal <- servidor + "no está disponible"
	} else {
		//fmt.Println(servidor, "está funcionando")
		canal <- servidor + "está funcionando"
	}
}

func main() {
	inicio := time.Now()

	canal := make(chan string)

	servidores := []string{
		"https://oregoom.com/",
		"https://www.udemy.com/",
		"https://www.facebook.com/",
		"https://www.google.com/",
	}

	for _, servidor := range servidores {
		go revisarServidor(servidor, canal)
	}

	for i := 0; i < len(servidores); i++ {
		fmt.Println(<-canal)
	}

	tiempoPaso := time.Since(inicio)
	fmt.Println("Tiempo de ejecución:", tiempoPaso)
}
