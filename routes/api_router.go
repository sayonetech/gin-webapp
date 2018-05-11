package routes

import (
	c "go-webapp/controller"

	"github.com/gin-gonic/gin"
)

func registerAPIRouter(router *gin.Engine) {

	api := router.Group("/api")
	api.GET("/index", c.IndexApi)

}
