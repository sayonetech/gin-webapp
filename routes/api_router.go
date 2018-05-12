package routes

import (
	c "go-webapp/controller"
	"go-webapp/controller/auth"

	"github.com/gin-gonic/gin"
)

func registerAPIRouter(router *gin.Engine) {

	api := router.Group("/api")
	api.GET("/index", c.IndexApi)

	v1 := router.Group("api/v1")
	{
		v1.POST("/register", auth.Register)
	}
}
