package router

import (
	"go-webapp/config"
	"go-webapp/handle"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	// proxy "github.com/chenhg5/gin-reverseproxy"
)

func initRouter() *gin.Engine {
	router := gin.New()

	if config.GetEnv().DEBUG {
		router.Use(gin.Logger()) // Used in development mode, console print request records
		pprof.Register(router)   // Performance Analysis Tool
	}

	router.Use(handle.HandleErrors()) // Error handling

	registerAPIRouter(router)

	// ReverseProxy
	// router.Use(proxy.ReverseProxy(map[string] string {
	// 	"localhost:4000" : "localhost:9090",
	// }))

	return router
}
