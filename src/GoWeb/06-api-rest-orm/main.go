package main

import (
	"fmt"
	"gorm/handlers"
	"gorm/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Migrate table Users
	models.MigrarUsers()
	// Rutas

	mux := mux.NewRouter()

	// EndPoint
	mux.HandleFunc("/api/user/", handlers.GetUsers).Methods("GET")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	mux.HandleFunc("/api/user/", handlers.CreateUser).Methods("POST")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")
	// Crear servidor
	fmt.Println("El servidor está corriendo en el puerto 8080")
	fmt.Println("Run server: http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
