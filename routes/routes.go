package routes

import (
	"log"
	"net/http"

	"github.com/Gabriel-Newton-dev/API_Rest_Golang/controllers"
	"github.com/gorilla/mux"
)

// HandleRequest(lhe dar com requisição) convenção de utilizar
func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home) // handleFunc - lhe dar com a função. // toda vez que ele receber um "/"" ele irá executar a nossa função Home.
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	log.Fatal(http.ListenAndServe(":7500", r))
}
