package router

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"lemon/app/controller/sd"
	"lemon/app/controller/v1/user"
	"lemon/app/middleware"
	"net/http"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())

	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(middleware.RequestId())

	g.Use(mw...)

	// 404 handler
	g.NoRoute(func(c *gin.Context) { c.String(http.StatusNotFound, "Not found.") })

	// 性能分析工具
	pprof.Register(g)

	// 系统健康检查
	svcd := g.Group("/api/sd")
	{
		svcd.GET("/health", sd.HealthCheck).GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
		svcd.GET("/host", sd.HostCheck)
		svcd.GET("/io", sd.IOCheck)
	}

	v1Group := g.Group("/api/v1")
	v1Group.GET("/test", user.Test)

	// web页面路由
	v1Group.POST("/login", user.Login)

	v1Group.Use(middleware.AuthMiddleware())
	{
		v1Group.GET("/user/:id", user.GetUserById)
		v1Group.GET("/users", user.GetUserList)
		v1Group.PUT("/user/:id", user.UpdateUserById)
		v1Group.DELETE("/user/:id", user.DeleteUserById)
	}

	return g
}
