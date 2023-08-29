package router

import (
	"mxshop_api/order_web/api/order"
	"mxshop_api/order_web/api/pay"
	"mxshop_api/order_web/middlewares"

	"github.com/gin-gonic/gin"
)

func InitOrderRouter(Router *gin.RouterGroup) {
	BannerRouter := Router.Group("orders").Use(middlewares.JWTAuth())
	{
		BannerRouter.GET("", order.List)       // 订单列表
		BannerRouter.POST("", order.New)       // 新建订单
		BannerRouter.GET("/:id", order.Detail) // 订单详情
	}
	PayRouter := Router.Group("pay")
	{
		PayRouter.POST("alipay/notify", pay.Notify)
	}
}
