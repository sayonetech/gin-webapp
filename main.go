package main

import (
	"go-webapp/config"
	"net/http"
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

	//Gin uses a custom version of HttpRouterhttps://github.com/julienschmidt/httprouter
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
