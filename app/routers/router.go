package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/varandrew/gin-product/app/controllers"
	"github.com/varandrew/gin-product/app/controllers/v1"
	"github.com/varandrew/gin-product/app/middlewares"
	"github.com/varandrew/gin-product/app/pkg/sd"
	"github.com/varandrew/gin-product/app/pkg/setting"
	_ "github.com/varandrew/gin-product/docs" // To resolve failed to load spec.
)

// @title Swagger Blog API
// @version 1.0
// @description This is a gin blog server.
// @termsOfService https://razeen.me

// @contact.name varandrew
// @contact.url https://github.com/varandrew
// @contact.email varandrewchen@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8000
// @BasePath /api/v1

func InitRouter() *gin.Engine {
	g := gin.New()

	g.Use(gin.Logger())
	g.Use(gin.Recovery()) // 在处理某些请求时可能因为程序 bug 或者其他异常情况导致程序 panic，这时候为了不影响下一次请求的调用，需要通过 gin.Recovery()来恢复 API 服务器

	g.Use(middlewares.NoCache) // 强制浏览器不使用缓存
	g.Use(middlewares.Options) // 浏览器跨域 OPTIONS 请求设置
	g.Use(middlewares.Secure)  // 一些安全设置

	gin.SetMode(setting.RunMode)

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	g.GET("/auth", controllers.GetAuth)

	apiV1 := g.Group("/api/v1")
	apiV1.Use(middlewares.JWT())
	{
		apiV1.GET("/tags", v1.GetTags)
		apiV1.POST("/tags", v1.AddTag)
		apiV1.PUT("/tags/:id", v1.EditTag)
		apiV1.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章列表
		apiV1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiV1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiV1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiV1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiV1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	g.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "test",
		})
	})

	return g
}
