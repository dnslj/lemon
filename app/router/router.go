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
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Not found.")
	})

	// pprof router
	pprof.Register(g)

	g.POST("/api/v1/login", user.Login)

	userGroup := g.Group("/api/v1/user")
	userGroup.Use(middleware.AuthMiddleware())
	{
		userGroup.GET("/:id", user.GetUserById)
		userGroup.GET("", user.GetUserList)
		userGroup.PUT("/:id", user.UpdateUserById)
		userGroup.DELETE("/:id", user.DeleteUserById)
	}

	// The health check handlers
	svcd := g.Group("/api/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
		svcd.GET("/host", sd.HostCheck)
		svcd.GET("/io", sd.IOCheck)
	}

	return g
}
