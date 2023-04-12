package routes

import (
	"github.com/gin-gonic/gin"
)

func SuperRouter() *gin.Engine {

	router := gin.Default()
	PingRoutes(router)
	PongRoutes(router)

	return router
}
