package routers

import (
	service "payment/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	account := router.Group("/account")
	{
		account.GET("/information", service.GetAccountInformation)
		account.POST("/create", service.CreateAccount)
	}
}
