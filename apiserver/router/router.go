package router

import (
	"../handle/sd"
	"../handle/user"
	"./middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {

	//在处理某些请求时可能因为程序bug或者其他异常情况导致程序panic，
	// 这时候为了不影响下一次请求的调用需要通过gin.Recovery来恢复API服务器
	g.Use(gin.Recovery())
	//强制浏览器不使用缓存
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	// 404 Handle
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route")
	})

	g.POST("/login", user.Login)

	//user
	u := g.Group("/v1/user")
	u.POST("", user.Create)
	u.Use(middleware.AuthMiddleware())
	{
		u.DELETE("/:id", user.Delete)
		u.PUT("/:id", user.Update)
		u.GET("", user.List)
		u.GET("/:username", user.Get)
	}

	//The health check handles
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g

}
