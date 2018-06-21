package routes

import (
	c "go-webapp/controller"
	"go-webapp/controller/auth"
	"go-webapp/middleware/session"
	"go-webapp/module/debug"
	"go-webapp/module/server"

	"github.com/gin-gonic/gin"
)

var store *session.Store

func init() {
	store = session.NewSessionStore()
}

func registerAPIRouter(router *gin.Engine) {

	api := router.Group("/api")
	api.GET("/index", c.IndexApi)

	v1 := router.Group("api/v1")
	{
		v1.POST("/register", auth.Register)
		v1.POST("/login", auth.UserLogin)
	}

	debugger := router.Group("/api/debug")
	{
		//TODO Session Must Admin
		debugger.GET("/pprof/", debug.IndexHandler())
		debugger.GET("/pprof/heap", debug.HeapHandler())
		debugger.GET("/pprof/goroutine", debug.GoroutineHandler())
		debugger.GET("/pprof/block", debug.BlockHandler())
		debugger.GET("/pprof/threadcreate", debug.ThreadCreateHandler())
		debugger.GET("/pprof/cmdline", debug.CmdlineHandler())
		debugger.GET("/pprof/profile", debug.ProfileHandler())
		debugger.GET("/pprof/symbol", debug.SymbolHandler())
		debugger.POST("/pprof/symbol", debug.SymbolHandler())
		debugger.GET("/pprof/trace", debug.TraceHandler())
	}

	router.GET("/version", server.Version)

}
