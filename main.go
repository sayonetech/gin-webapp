package main

import (
	"go-webapp/config"
	"go-webapp/module/server"
	"go-webapp/routes"

	"runtime"

	"github.com/gin-gonic/gin"
)

/**
* run main.go(go run main.go) and visit 0.0.0.0:8080/ping on browser
 */
func main() {

	// Set maximum number of CPUs that can be executing simultaneously with the number of logical CPUs usable by the current process
	runtime.GOMAXPROCS(runtime.NumCPU())

	if config.GetEnv().DEBUG {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := routes.InitRouter() // 初始化路由

	server.Run(router)

}
