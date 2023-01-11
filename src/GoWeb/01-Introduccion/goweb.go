package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handlers
func Hola(response http.ResponseWriter, request *http.Request) {
	fmt.Println("El método es", request.Method)
	fmt.Fprintln(response, "Hola Mundo")
}

func PageNotFound(response http.ResponseWriter, request *http.Request) {
	http.NotFound(response, request)
}

func Error(response http.ResponseWriter, request *http.Request) {
	http.Error(response, "La página tuvo un error", http.StatusNotFound)
}

func Saludar(response http.ResponseWriter, request *http.Request) {
	fmt.Println(request.URL)
	fmt.Println(request.URL.Query())
	name := request.URL.Query().Get("name")
	age := request.URL.Query().Get("age")
	fmt.Fprintf(response, "Hola %s tu edad es %s", name, age)
}
func main() {
	// Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/page", PageNotFound)
	mux.HandleFunc("/error", Error)
	mux.HandleFunc("/saludar", Saludar)

	// Router
	//http.HandleFunc("/", Hola)
	//http.HandleFunc("/page", PageNotFound)
	//http.HandleFunc("/error", Error)
	//http.HandleFunc("/saludar", Saludar)

	// Crear servidor
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	fmt.Println("El servidor está corriendo en el puerto 8080")
	fmt.Println("Run server: http://localhost:8080/")
	//log.Fatal(http.ListenAndServe("localhost:8080", mux))
	log.Fatal(server.ListenAndServe())
}
