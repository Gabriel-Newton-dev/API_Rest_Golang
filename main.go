package main

import (
	"fmt"

	"github.com/Gabriel-Newton-dev/API_Rest_Golang/database"
	"github.com/Gabriel-Newton-dev/API_Rest_Golang/routes"
	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando nosso servidor Rest em Go")
	routes.HandleRequest()

}
