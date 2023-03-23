package routes

import (
	"blog/controllers"
	"blog/logger"
	"blog/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1")
	// 注册业务路由
	v1.POST("/signup", controllers.SignUpHandler)
	v1.POST("/login", controllers.LoginHandler)
	v1.Use(middlewares.JWTAuthMiddleware()) // 应用JWT认证中间件
	{
		v1.GET("/community", controllers.CommunityHandler)
		v1.GET("/community/:id", controllers.CommunityDetailHandler)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
