package routes

import (
	"firstDevProject/controllers"

	"github.com/gin-gonic/gin"
)

func PingRouter(router *gin.Engine) {

	router.GET("/ping", controllers.Ping)

}
