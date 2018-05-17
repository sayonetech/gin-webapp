package routes

import (
	c "go-webapp/controller"
	"go-webapp/controller/auth"
	"go-webapp/middleware/session"

	"github.com/gin-gonic/gin"
)

func registerAPIRouter(router *gin.Engine) {
	router.Use(session.SessionMiddleWare())
	api := router.Group("/api")
	api.GET("/index", c.IndexApi)

	v1 := router.Group("api/v1")
	{
		v1.POST("/register", auth.Register)
		v1.POST("/login", auth.UserLogin)
	}

}
