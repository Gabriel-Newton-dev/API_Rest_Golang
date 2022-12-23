package routes

import (
	"log"
	"net/http"

	"github.com/Gabriel-Newton-dev/API_Rest_Golang/controllers"
)

// HandleRequest(lhe dar com requisição) convenção de utilizar
func HandleRequest() {
	http.HandleFunc("/", controllers.Home) // handleFunc - lhe dar com a função.
	// toda vez que ele receber um "/"" ele irá executar a nossa função Home.
	log.Fatal(http.ListenAndServe(":7500", nil))
}
