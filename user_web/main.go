package main

import (
	"fmt"
	"mxshop_api/user_web/global"
	"mxshop_api/user_web/initialize"
	"mxshop_api/user_web/utils"
	myvalidator "mxshop_api/user_web/validator"

	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {

	initialize.InitLogger()
	initialize.InitConfig()
	Router := initialize.Routers()
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}
	initialize.InitSrvConn()

	viper.AutomaticEnv()
	// 如果是本地开发环境，端口号需要固定，线上环境启动获取端口号
	debug := viper.GetBool("MXSHOP_DEBUG")
	if !debug {
		port, err := utils.GetFreePort()
		if err == nil {
			global.ServerConfig.Port = port
		}
	}

	// 注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", myvalidator.ValidateMobile)
		_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 非法的手机号码！", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}
	// defer logger.Sync()
	// suger := logger.Sugar()

	// port := 8021
	zap.S().Debugf("启动服务器，端口：%d", global.ServerConfig.Port)

	if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("启动失败：", err.Error())
	}
}
