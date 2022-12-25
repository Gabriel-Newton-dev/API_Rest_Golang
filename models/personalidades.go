package models

type Personalidade struct {
	Id       int
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
}

// Slice de Personalidade dentro da variavel Personalidades.
var Personalidades []Personalidade
