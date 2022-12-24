package models

type Personalidade struct {
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
}

// Slice de Personalidade dentro da variavel Personalidades.
var Personalidades []Personalidade
