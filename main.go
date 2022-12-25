package main

import (
	"fmt"

	"github.com/Gabriel-Newton-dev/API_Rest_Golang/models"
	"github.com/Gabriel-Newton-dev/API_Rest_Golang/routes"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{1, "Carmela Dutra", "Carmela Teles Leite Dutra (Rio de Janeiro, 17 de setembro de 1884 — Rio de Janeiro, 9 de outubro de 1947) foi a primeira-dama do Brasil, de 31 de janeiro de 1946 até a sua morte, tendo sido a esposa de Eurico Gaspar Dutra, 16.º Presidente do Brasil. Era, carinhosamente, chamada de Dona Santinha, pela sua forte religiosidade, fazendo seu marido abrir uma capelinha no Palácio Guanabara."},
		{2, "Princesa izabel", "Isabel Cristina (Rio de Janeiro, 29 de julho de 1846 – Eu, 14 de novembro de 1921), cognominada a Redentora, foi a filha mais velha do imperador Pedro II do Brasil e da imperatriz consorte Teresa Cristina das Duas Sicílias e, portanto membro do ramo brasileiro da Casa de Bragança. Como a Herdeira presuntiva do Império do Brasil, ela recebeu o título de Princesa Imperial e foi Regente do Império em três ocasiões diferentes — numa delas, assinou a Lei Áurea, que declarou extinta a escravidão no Brasil."},
	}

	fmt.Println("Iniciando nosso servidor Rest em Go")
	routes.HandleRequest()

}
