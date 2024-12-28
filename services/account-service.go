package service

import (
	"fmt"
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

func CreateAccount(c *gin.Context) {
	db := databases.GetWriteInstance()
	var account models.Account
	fmt.Println(c.Request.Body)
	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	tx := db.Begin()
	if tx.Error != nil {
		c.JSON(500, gin.H{"error": tx.Error.Error()})
		return
	}

	if result := tx.Create(&account); result.Error != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	// Assuming AccountNumberAddress is a related model that needs to be created
	var accountNumberAddress models.AccountAddressInformation
	accountNumberAddress.AccountID = account.ID
	if result := tx.Create(&accountNumberAddress); result.Error != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, account)
}
