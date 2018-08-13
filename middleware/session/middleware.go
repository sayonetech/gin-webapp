package session

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/context"
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

func Sessions(store *Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("store", store)
		defer context.Clear(c.Request)
		c.Next()
	}
}

func User(c *gin.Context) Store {
	//TODO
	return FromContext(c)
}

func MustUser() gin.HandlerFunc {
	//TODO
	return func(c *gin.Context) {
		//user := User(c)
		/*switch {
		case user == nil:
			c.String(401, "User not authorized")
			c.Abort()
		default:
			c.Next()
		}*/
	}
}
