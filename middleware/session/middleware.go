package session

import (
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
}
