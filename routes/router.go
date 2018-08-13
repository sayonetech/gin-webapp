package routes

import (
	"go-webapp/config"
	"go-webapp/middleware/cors"
	"go-webapp/middleware/log"
	"go-webapp/middleware/request"
	"go-webapp/middleware/session"

	"github.com/gin-gonic/gin"
	//	"os"
	// proxy "github.com/chenhg5/gin-reverseproxy"
	"github.com/getsentry/raven-go"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//InitRouter Initialise router
func InitRouter() *gin.Engine {

	raven.SetDSN(config.GetEnv().SENTRY_URL)
	route := gin.New()
	//reads all environment variables and sets them to a map
	//route.Use(gzip.Gzip(gzip.DefaultCompression))
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//var logger = loge.New()
	//	Log := log.Logger{Writer: logger, ErrorLog: os.Stdout}
	//route.Use(Log.JSONLogMiddleware())
	//TODO: open hte files to whcih to write log files
	//TODO: initialte the logger struct here and call the recovery function as a method
	route.Use(log.Recovery()) // *custom recovery
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
