package service

import (
	"net/http"
	databases "payment/database"
	"payment/models"

	"github.com/gin-gonic/gin"
)

func CreateExternalBank(c *gin.Context) {
	// create external bank
	db := databases.GetWriteInstance()
	var externalBank models.ExternalBank
	if err := c.ShouldBindJSON(&externalBank); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := db.Create(&externalBank); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save external bank"})
		return
	}
}
