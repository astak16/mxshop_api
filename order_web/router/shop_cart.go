package router

import (
	"mxshop_api/order_web/api/shop_cart"
	"mxshop_api/order_web/middlewares"

	"github.com/gin-gonic/gin"
)

func InitShopCartRouter(router *gin.RouterGroup) {
	GoodsRouter := router.Group("shopcarts").Use(middlewares.JWTAuth())

	{
		GoodsRouter.GET("", shop_cart.List)          // 购物车列表
		GoodsRouter.DELETE("/:id", shop_cart.Delete) // 删除条目
		GoodsRouter.POST("", shop_cart.New)          // 添加商品到购物车
		GoodsRouter.PATCH("/:id", shop_cart.Update)  // 修改条目

	}
}
