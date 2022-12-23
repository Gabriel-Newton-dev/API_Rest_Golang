package main

import (
	"fmt"
)

func main() {
	fmt.Println("Iniciando nosso servidor Rest em Go")
	controllers.HandleRequest()

}
