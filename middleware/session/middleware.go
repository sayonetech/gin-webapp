package session

import (
	"go-webapp/models"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func SessionMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		cookies := context.Request.Cookies()
		log.WithFields(log.Fields{
			"cookies": cookies,
		}).Info("Reading cookies")
	}

	//https://stackoverflow.com/questions/47085046/gin-sessions-stores-the-status-and-the-code-in-the-url-i-want-to-change-that-t
	//https://stackoverflow.com/questions/36122999/how-to-create-a-authentication-model-to-restful-api-using-golang-gin
	//https://github.com/Depado/gin-auth-example/blob/master/main.go
	//https://sosedoff.com/2014/12/21/gin-middleware.html
	//https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin
	//https://github.com/rageix/ginAuth/blob/master/auth.go
	//https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin
}

func User(c *gin.Context) *models.User {
	v, ok := c.Get("user")
	if !ok {
		return nil
	}
	u, ok := v.(*models.User)
	if !ok {
		return nil
	}
	return u
}

func MustUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := User(c)
		switch {
		case user == nil:
			c.String(401, "User not authorized")
			c.Abort()
		default:
			c.Next()
		}
	}
}
