package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Gabriel-Newton-dev/API_Rest_Golang/models"
)

// Home toda vez que chegar uma requisicao Home, a func irá escrever writer (w), a String que colocamos que no caso é "Home Page"
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")

}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Personalidade)
}
