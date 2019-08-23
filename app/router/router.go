package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/gin-contrib/pprof"
	"study/lemon/app/middleware"
	"study/lemon/app/controller/v1/user"
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

	return g
}
