package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Gabriel-Newton-dev/API_Rest_Golang/models"
	"github.com/gorilla/mux"
)

// Home toda vez que chegar uma requisicao Home, a func irá escrever writer (w), a String que colocamos que no caso é "Home Page"
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")

}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidades)
}

func RetornaUmaPersoalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // aqui nos temos que colocar qual a request que estamos recebendo, como padrao o r r *http.Request
	id := vars["id"]

	for _, personalidade := range models.Personalidades {
		if strconv.Itoa(personalidade.Id) == id {
			json.NewEncoder(w).Encode(personalidade)
		}
	}
}
