package main

import (
	"github.com/Gabriel-Newton-dev/API_Rest_Golang/routes"
)

func main() {

	// viper.SetConfigFile(".env")
	// viper.ReadInConfig()
	// database.ConectaComBancoDeDados()
	// fmt.Println("Iniciando nosso servidor Rest em Go")
	routes.HandleRequest()

}
