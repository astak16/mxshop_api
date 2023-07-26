package initialize

import (
	"mxshop_api/user_web/global"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}
func InitConfig() {
	debug := GetEnvInfo("MXSHOP_DEBUG")
	if debug {
		viper.SetConfigName("config-debug")
	} else {
		viper.SetConfigName("config-pro")
	}

	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&global.ServerConfig); err != nil {
		panic(err)
	}

	zap.S().Infof("配置信息: %v", global.ServerConfig)
	viper.OnConfigChange(func(e fsnotify.Event) {
		zap.S().Infof("配置文件产生变化: %s", e.Name)
		_ = viper.ReadInConfig()
		_ = viper.Unmarshal(&global.ServerConfig)
		zap.S().Infof("配置信息: %v", global.ServerConfig)
	})
	viper.WatchConfig()
}
