package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Gabriel-Newton-dev/API_Rest_Golang/database"
	"github.com/Gabriel-Newton-dev/API_Rest_Golang/models"
)

// Home toda vez que chegar uma requisicao Home, a func irá escrever writer (w), a String que colocamos que no caso é "Home Page"
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")

}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var personalidades []models.Personalidade
	database.DB.Find(&personalidades)
	json.NewEncoder(w).Encode(personalidades)
}

// func RetornaUmaPersoalidade(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r) // aqui nos temos que colocar qual a request que estamos recebendo, como padrao o r r *http.Request
// 	id := vars["id"]

// 	for _, personalidade := range models.Personalidades {
// 		if strconv.Itoa(personalidade.Id) == id {
// 			json.NewEncoder(w).Encode(personalidade)
// 		}
// 	}
// }
