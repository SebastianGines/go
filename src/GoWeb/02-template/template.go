package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// Funciones
func Saludar(name string) string {
	return "Hola " + name + " desde una función"
}

// Estructuras
type Usuarios struct {
	UserName string
	Age      int
	Active   bool
	Admin    bool
	Cursos   []Curso
}

type Curso struct {
	Name string
}

// Variables globales
// Map de funciones
var funciones = template.FuncMap{
	"saludar": Saludar,
}

// ParseGlob para que guarde todos los templates htmls de una carpeta
var templates = template.Must(template.New("T").Funcs(funciones).ParseGlob("templates/**/*.html"))
var errorTemplate = template.Must(template.ParseFiles("templates/error/error.html"))

// Handler Error
func handleError(response http.ResponseWriter, status int) {
	response.WriteHeader(status)
	errorTemplate.Execute(response, nil)
}

// Función de render template
func renderTemplate(response http.ResponseWriter, name string, data interface{}) {
	err := templates.ExecuteTemplate(response, name, data)
	if err != nil {
		handleError(response, http.StatusInternalServerError)
	}
}

// Handler
func Index(response http.ResponseWriter, request *http.Request) {
	curso1 := Curso{"Go"}
	curso2 := Curso{"Python"}
	curso3 := Curso{"Java"}
	curso4 := Curso{"JavaScript"}
	usuario := Usuarios{"Seba", 35, true, false, []Curso{curso1, curso2, curso3, curso4}}
	renderTemplate(response, "index.html", usuario)
	// Template con errores
	/*
		template, error := template.New("index.html").Funcs(funciones).
			ParseFiles("index.html", "base.html")

		if error != nil {
			panic(error)
		} else {
			template.Execute(response, usuario)
		}*/

	// Template sin manejo de errores y utiliza Must
	/*
		template := template.Must(template.New("index.html").Funcs(funciones).
			ParseFiles("index.html", "base.html"))
			template.Execute(response, usuario)
	*/
}

func Registro(response http.ResponseWriter, request *http.Request) {
	renderTemplate(response, "registro.html", nil)
}

func main() {
	// Mux
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/registro", Registro)

	// Crear servidor
	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	fmt.Println("El servidor está corriendo en el puerto 8080")
	fmt.Println("Run server: http://localhost:8080/")
	log.Fatal(server.ListenAndServe())
}
