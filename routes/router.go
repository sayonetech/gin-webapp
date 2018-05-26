package routes

import (
	"go-webapp/config"
	"go-webapp/handle"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	// proxy "github.com/chenhg5/gin-reverseproxy"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//InitRouter Initialise router
func InitRouter() *gin.Engine {
	route := gin.New()
	store, _ := redis.NewStore(10, "tcp", config.GetEnv().REDIS_HOST, "", []byte("secret"))
	route.Use(sessions.Sessions(config.GetEnv().SESSION_OBJ_KEY, store))
	//route.Use(gzip.Gzip(gzip.DefaultCompression))
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	if config.GetEnv().DEBUG {
		route.Use(gin.Logger()) // Used in development mode, console print request records
	}

	route.Use(handle.Errors()) // Error handling

	registerAPIRouter(route)

	// ReverseProxy
	// router.Use(proxy.ReverseProxy(map[string] string {
	// 	"localhost:4000" : "localhost:9090",
	// }))

	return route
}
