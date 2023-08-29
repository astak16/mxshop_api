package initialize

import (
	"mxshop_api/userop_web/middlewares"
	"mxshop_api/userop_web/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"code": 200, "success": true})
	})

	// 配置跨域
	Router.Use(middlewares.Cors())
	ApiGroup := Router.Group("/up/v1")
	router.InitAddressRouter(ApiGroup)
	router.InitUserFavRouter(ApiGroup)
	router.InitMessageRouter(ApiGroup)

	return Router
}
