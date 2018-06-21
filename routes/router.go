package routes

import (
	"github.com/gin-gonic/gin"
	loge "github.com/sirupsen/logrus"
	"go-webapp/config"
	//	"go-webapp/handle"
	"go-webapp/middleware"
	"go-webapp/middleware/cors"
	"go-webapp/middleware/log"
	"go-webapp/middleware/request"
	"go-webapp/middleware/session"
	"os"
	// proxy "github.com/chenhg5/gin-reverseproxy"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//InitRouter Initialise router
func InitRouter() *gin.Engine {
	route := gin.New()
	//route.Use(gzip.Gzip(gzip.DefaultCompression))
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	var errorLog = loge.New()

	if config.GetEnv().DEBUG {
		//		route.Use(gin.Logger()) // Used in development mode, console print request records
		loge.SetOutput(os.Stdout) //set output to console
	} else {
		file, err := os.OpenFile(config.GetEnv().ERROR_LOG_PATH, os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			panic(err) //if log files are not present server will panic
		}
		errorLog.Out = file
	}
	route.Use(log.JSONLogMiddleware())
	//route.Use(gin.Recovery())
	route.Use(middleware.Recovery()) // *custom recovery
	route.Use(request.RequestID(request.RequestIDOptions{AllowSetting: false}))
	route.Use(cors.CORS(cors.CORSOptions{}))
	//route.Use(handle.Errors()) // Error handling
	route.Use(session.Sessions(store))
	route.Use(session.SessionMiddleWare())
	registerAPIRouter(route)

	// ReverseProxy
	// router.Use(proxy.ReverseProxy(map[string] string {
	// 	"localhost:4000" : "localhost:9090",
	// }))

	return route
}
