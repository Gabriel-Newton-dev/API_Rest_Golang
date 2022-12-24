package main

import (
	"fmt"
	"log"
	"net/http"
)

// Home toda vez que chegar uma requisicao Home, a func irá escrever writer (w), a String que colocamos que no caso é "Home Page"
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")

}

// HandleRequest(lhe dar com requisição) convenção de utilizar
func HandleRequest() {
	http.HandleFunc("/", Home) // handleFunc - lhe dar com a função.
	// toda vez que ele receber um "/"" ele irá executar a nossa função Home.
	log.Fatal(http.ListenAndServe(":7500", nil))
}

func main() {
	fmt.Println("Iniciando nosso servidor Rest em Go")
	HandleRequest()

}
