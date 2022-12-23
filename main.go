package main

import (
	"encoding/json"
	"fmt"
)

type Cadastro struct {
	Nome     string
	Idade    int
	Endereco string
	Saldo    float64
}

func main() {

	Api := "Api rest com Golang"
	fmt.Println(Api)

	Gabriel := Cadastro{"Gabriel", 36, "Miami", 645.67}

	GabrielJson, err := json.Marshal(Gabriel)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(GabrielJson))

}
