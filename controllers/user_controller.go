package controllers

import (
	"github.com/Gabriel-Newton-dev/API_Rest_Golang/database"
	"github.com/Gabriel-Newton-dev/API_Rest_Golang/models"
	"github.com/gin-gonic/gin"
)

// cria uma funcao de criar usuário

func CreateUser(c *gin.Context) {
	db := database.GetDatabase()

	var u models.User

	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON:" + err.Error(),
		})
		return
	}

	err := db.Create(&u).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create user" + err.Error(),
		})
		return
	}
	c.JSON(200, u)

}
