package routes

import (
	"firstDevProject/controllers"

	"github.com/gin-gonic/gin"
)

func PongRoutes(router *gin.Engine) {

	router.GET("/pong/:pingORPong", controllers.Pong)
}
