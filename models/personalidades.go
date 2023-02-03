package models

import "gorm.io/gorm"

type Personalidade struct {
	gorm.Model
	Id       int    `json:"id"`
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
}

// Slice de Personalidade dentro da variavel Personalidades.
var Personalidades []Personalidade
