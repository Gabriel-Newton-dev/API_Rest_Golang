package main

import (
	"fmt"

	"github.com/Gabriel-Newton-dev/API_Rest_Golang/routes"
)

func main() {

	models.Personalidades

	fmt.Println("Iniciando nosso servidor Rest em Go")
	routes.HandleRequest()

}
