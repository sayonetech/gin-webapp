package middleware

import (
	"github.com/gin-gonic/gin"
	loge "github.com/sirupsen/logrus"
	"go-webapp/common"
	"runtime/debug"
)

func Recovery() gin.HandlerFunc {

	return func(c *gin.Context) {

		defer func() {
			if recover() != nil {
				c.JSON(500, gin.H{
					"code": 10500,
					"msg":  "internal",
				})

				loge.WithFields(loge.Fields{
					"client_ip":  common.GetClientIP(c),
					"method":     c.Request.Method,
					"path":       c.Request.RequestURI,
					"status":     c.Writer.Status(),
					"referrer":   c.Request.Referer(),
					"request_id": c.Writer.Header().Get("Request-Id"),
					"stack":      string(debug.Stack()),
					// "api_version": util.ApiVersion,
				}).Error(c.Errors.String())
			}
		}()
		// resume by calling gin context next
		c.Next()
	}

}
