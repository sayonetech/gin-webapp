package log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-webapp/common"
	"go-webapp/config"
	"io"
	"os"
	"runtime/debug"
	"time"
)

type Logger struct {
	//	AccessLog io.Writer // pointer to Access Log
	ErrorLog io.Writer // pointer to Access Log
	//	AccessLog io.Writer // pointer to Access Log
	Writer *log.Logger
}

// JSONLogMiddleware logs a gin HTTP request in JSON format, with some additional custom key/values
func (l *Logger) JSONLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Process Request
		//l.Writer.Out = l.AccessLog
		if c.Writer.Status() == 200 {
			l.Log(c).Info("")
		} else {
			l.Log(c).Error(c.Errors.String())
		}
		c.Next()
	}
}

func Recovery() gin.HandlerFunc {

	return func(c *gin.Context) {

		defer func() {

			if recover() != nil {

				errorFile, _ := os.OpenFile(config.GetEnv().ERROR_LOG_PATH, os.O_CREATE|os.O_WRONLY, 0666)
				c.JSON(500, gin.H{
					"code": 10500,
					"msg":  "internal",
				})

				if config.GetEnv().DEBUG {

				} else {
					log.SetOutput(io.MultiWriter(errorFile))
				}

				log.WithFields(log.Fields{
					"client_ip":  common.GetClientIP(c),
					"method":     c.Request.Method,
					"path":       c.Request.RequestURI,
					"status":     c.Writer.Status(),
					"referrer":   c.Request.Referer(),
					"request_id": c.Writer.Header().Get("Request-Id"),
					"stack":      string(debug.Stack()),
					// "api_version": util.ApiVersion,
				}).Error(c.Errors.String())
				_ = errorFile.Close()
				fmt.Println("...")
			}
		}()
		// resume by calling gin context next
		c.Next()
	}

}

func (l *Logger) Log(c *gin.Context) *log.Entry {

	// Start timer
	start := time.Now()

	// Stop timer
	duration := common.GetDurationInMillseconds(start)

	return l.Writer.WithFields(log.Fields{
		"client_ip":  common.GetClientIP(c),
		"duration":   duration,
		"method":     c.Request.Method,
		"path":       c.Request.RequestURI,
		"status":     c.Writer.Status(),
		"referrer":   c.Request.Referer(),
		"request_id": c.Writer.Header().Get("Request-Id"),
		// "api_version": util.ApiVersion,
	})
}
