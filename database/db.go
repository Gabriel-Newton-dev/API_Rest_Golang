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

	DB_USER := viper.Get("DATABASE_USER")
	DB_NAME := viper.Get("DATABASE_NAME")
	DB_PASSWORD := viper.Get("DATABASE_PASSWORD")

	stringDeConexao := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	DB, err := gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados.")
	}
	DB.AutoMigrate(&models.Personalidade{})
}

// Caso precise utilziar um banco de dados a parte sem ser pelo docker.
