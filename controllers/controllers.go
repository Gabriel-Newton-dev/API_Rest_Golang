package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Gabriel-Newton-dev/API_Rest_Golang/database"
	"github.com/Gabriel-Newton-dev/API_Rest_Golang/models"
	"github.com/gorilla/mux"
)

var ID int

// Home toda vez que chegar uma requisicao Home, a func irá escrever writer (w), a String que colocamos que no caso é "Home Page"
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")

}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}
