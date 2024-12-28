package service

import (
	databases "payment/database"
	"payment/models"

	"github.com/gin-gonic/gin"
)

func GetAccountInformation(c *gin.Context) {
	db := databases.GetWriteInstance()
	var accounts []models.Account
	if result := db.Find(&accounts); result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(200, accounts)
}

func UpdateAccountCredit(c *gin.Context) {}

func CreateAccount(c *gin.Context) {}
