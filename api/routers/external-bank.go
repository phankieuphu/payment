package routers

import (
	service "payment/services"

	"github.com/gin-gonic/gin"
)

func ExternalBankRouter(router *gin.Engine) {
	externalBankRouter := router.Group("/external-bank")
	{
		externalBankRouter.POST("/create-bank", service.CreateExternalBank)
	}
}
