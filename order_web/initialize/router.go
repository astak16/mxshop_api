package initialize

import (
	"mxshop_api/order_web/middlewares"
	"mxshop_api/order_web/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"code": 200, "success": true})
	})

	// 配置跨域
	Router.Use(middlewares.Cors())
	ApiGroup := Router.Group("/o/v1")
	router.InitOrderRouter(ApiGroup)
	router.InitShopCartRouter(ApiGroup)

	return Router
}
