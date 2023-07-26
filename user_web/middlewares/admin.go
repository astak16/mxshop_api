package middlewares

import (
	"fmt"
	"mxshop_api/user_web/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsAdminAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, _ := ctx.Get("claims")
		currentUser := claims.(*models.CustomClaims)
		fmt.Println(currentUser.AuthorityId)
		if currentUser.AuthorityId != 2 {
			ctx.JSON(http.StatusForbidden, gin.H{
				"msg": "无权限",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
