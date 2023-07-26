package router

import (
	"mxshop_api/user_web/api"
	"mxshop_api/user_web/middlewares"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InitUserRouter(router *gin.RouterGroup) {
	UserRouter := router.Group("/user")
	zap.S().Info("配置用户相关的 router")
	{
		UserRouter.GET("/list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
		UserRouter.POST("/pwd_login", api.PassWordLogin)
		UserRouter.POST("/register", api.Register)
	}
}
