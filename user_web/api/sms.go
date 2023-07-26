package api

import (
	"fmt"
	"math/rand"
	"mxshop_api/user_web/forms"
	"mxshop_api/user_web/global"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func GenerateSmsCode(width int) string {
	numeric := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

func SendSms(ctx *gin.Context) {
	// fmt.Println("SendSms")
	sendSmsForm := forms.SendSmsForm{}
	if err := ctx.ShouldBindJSON(&sendSmsForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}
	// fmt.Println(sendSmsForm)

	code := GenerateSmsCode(6)

	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port),
	})
	rdb.Set(ctx, sendSmsForm.Mobile, code, time.Duration(global.ServerConfig.AliSmsInfo.Expire)*time.Second)
	c := rdb.Get(ctx, sendSmsForm.Mobile)
	fmt.Println(c)
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "发送成功",
	})
}
