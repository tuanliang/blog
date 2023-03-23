package routes

import (
	"blog/controllers"
	"blog/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册业务路由
	r.POST("/signup", controllers.SignUpHandler)
	r.POST("/login", controllers.LoginHandler)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	return r
}
