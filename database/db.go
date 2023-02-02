package database

import (
	"fmt"
	"log"

	"github.com/Gabriel-Newton-dev/API_Rest_Golang/models"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {

	DATABASE_USER := viper.Get("DATABASE_USER")
	DATABASE_NAME := viper.Get("DATABASE_NAME")
	DATABASE_PASSWORD := viper.Get("DATABASE_PASSWORD")

	StringDeConexao := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s sslmode=disable", DATABASE_USER, DATABASE_PASSWORD, DATABASE_NAME)
	DB, err := gorm.Open(postgres.Open(StringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados.")
	}
	DB.AutoMigrate(&models.Personalidade{})
}

// Caso precise utilziar um banco de dados a parte sem ser pelo docker.
