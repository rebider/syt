package router

import (
	"github.com/gin-gonic/gin"
	"github.com/aveyuan/syt/controller"
	"github.com/aveyuan/syt/middleware"
)

func Init()*gin.Engine  {
	r := gin.Default()
	//注册中间件
	middleware.Init(r)
	//单独注册jwt中间件
	authMiddleware:=middleware.Jwtmiddleware(r)
	//把r给controller
	controller.R = r

	r.GET("/", func(c *gin.Context) {
		c.String(200,"hello world")
	})

	r.POST("/reg",controller.ClientRegPost)


	r.NoRoute(func(c *gin.Context) {
		controller.ResJson(404,"Page Not Found",c)
	})

	r.POST("/login", authMiddleware.LoginHandler)

	auth := r.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/home",controller.Home)
		auth.GET("/usertickets",controller.UserTickets)
		auth.GET("/listtk",controller.ListTickets)
	}


	return r
}
