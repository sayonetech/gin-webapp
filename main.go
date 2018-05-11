package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
* run main.go(go run main.go) and visit 0.0.0.0:8080/ping on browser
 */
func main() {
	//Gin uses a custom version of HttpRouterhttps://github.com/julienschmidt/httprouter
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
