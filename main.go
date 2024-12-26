package main

import (
	"payment/api/routers"
	databases "payment/database"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	databases.InitWriteDB()
	router := gin.Default()

	routers.SetupRouter(router)

	router.Run(":8080")
}
