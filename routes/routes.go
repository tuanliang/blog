package routes

import (
	"blog/controller"
	_ "blog/docs" // 千万不要忘了导入把你上一步生成的docs
	"blog/logger"
	"blog/middleware"
	"net/http"

	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	// r.Use(logger.GinLogger(), logger.GinRecovery(true), middleware.RateLimitMiddleware(2*time.Second, 1))
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.LoadHTMLFiles("./templates/index.html")
	r.Static("/static", "./static")
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ok")
	})

	v1 := r.Group("/api/v1")
	// 注册业务路由
	v1.POST("/signup", controller.SignUpHandler)
	v1.POST("/login", controller.LoginHandler)

	v1.Use(middleware.JWTAuthMiddleware()) // 应用jwt中间件
	v1.GET("/posts2", controller.GetPostListHandler2)

	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDetailHandler)
		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		v1.GET("/posts", controller.GetPostListHandler)

		v1.POST("/vote", controller.PostVoteController)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
