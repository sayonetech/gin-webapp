package server

import (
	"go-webapp/config"

	"github.com/gin-gonic/gin"
)

// Version is the specification version that the package types support.

// Health endpoint returns a 500 if the server state is unhealthy.
func Health(c *gin.Context) {
	//TODO Implement context
	/*
		if err := store.FromContext(c).Ping(); err != nil {
			c.String(500, err.Error())
			return
		}
		c.String(200, "")
	*/
}

// Version endpoint returns the server version and build information.
func Version(c *gin.Context) {
	c.JSON(200, gin.H{
		"source":  "https://github.com/sayonetech/gin-webapp",
		"version": config.GetEnv().VERSION,
	})
}
