package handlers

import (
	"apirest/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	//rw.Header().Set("Content-Type", "text/xml")
	users, _ := models.ListUsers()
	output, _ := json.Marshal(users)
	//output, _ := xml.Marshal(users)
	//output, _ := yaml.Marshal(users)
	fmt.Fprintln(rw, string(output))
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	// Obtener id
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])
	user, _ := models.GetUser(userId)
	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output))
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	// Obtener registro a crear
	user := models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		fmt.Fprintln(rw, http.StatusUnprocessableEntity)
	} else {
		user.Save()
	}
	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output))
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	// Obtener registro a crear
	user := models.User{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		fmt.Fprintln(rw, http.StatusUnprocessableEntity)
	} else {
		user.Save()
	}
	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output))
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	// Obtener id
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])
	user, _ := models.GetUser(userId)
	user.Delete()
	output, _ := json.Marshal(user)
	fmt.Fprintln(rw, string(output))
}
