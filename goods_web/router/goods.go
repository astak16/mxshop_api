package router

import (
	"mxshop_api/goods_web/api/goods"
	"mxshop_api/goods_web/middlewares"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitGoodsRouter(router *gin.RouterGroup) {
	GoodsRouter := router.Group("/goods")
	zap.S().Info("配置商品相关的 router")
	{
		GoodsRouter.GET("", goods.List)                                                                 // 商品列表
		GoodsRouter.POST("", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.New)               // 新增商品
		GoodsRouter.GET("/:id", goods.Detail)                                                           // 商品详情
		GoodsRouter.DELETE("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.Delete)      // 删除商品
		GoodsRouter.GET("/:id/stocks", goods.Stocks)                                                    // 商品库存
		GoodsRouter.PUT("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.Update)         // 更新商品信息
		GoodsRouter.PATCH("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.UpdateStatus) // 更新商品部分信息
	}
}
