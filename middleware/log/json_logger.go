package log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-webapp/common"
	"go-webapp/config"
	"io"
	//"os"
	"runtime/debug"
	"time"
)

// TODO: add the recover functiont to logger struct
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

				c.JSON(500, gin.H{
					"code": 10500,
					"msg":  "internal",
				})
				if config.GetEnv().DEBUG {
					//log.SetOutput(os.Stdout)
					fmt.Println(fmt.Errorf("%s", string(debug.Stack())))
				} else {
					CaptureErrorWithSentry(fmt.Errorf("%s", string(debug.Stack())))
				}
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
