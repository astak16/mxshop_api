package initialize

import (
	"encoding/json"
	"mxshop_api/userop_web/global"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
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

	if err := viper.Unmarshal(global.NacosConfig); err != nil {
		panic(err)
	}

	// zap.S().Infof("配置信息: %v", global.NacosConfig)
	// viper.OnConfigChange(func(e fsnotify.Event) {
	// 	zap.S().Infof("配置文件产生变化: %s", e.Name)
	// 	_ = viper.ReadInConfig()
	// 	_ = viper.Unmarshal(&global.NacosConfig)
	// 	zap.S().Infof("配置信息: %v", global.NacosConfig)
	// })
	// viper.WatchConfig()

	sc := []constant.ServerConfig{
		{
			IpAddr: global.NacosConfig.Host,
			Port:   global.NacosConfig.Port,
		},
	}
	cc := constant.ClientConfig{
		NamespaceId:         global.NacosConfig.Namespace,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "../tmp/nacos/log",
		CacheDir:            "../tmp/nacos/cache",
		LogLevel:            "debug",
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})

	if err != nil {
		panic(err)
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: global.NacosConfig.DataId,
		Group:  global.NacosConfig.Group,
	})

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(content), &global.ServerConfig)

	if err != nil {
		zap.S().Fatalf("读取nacos配置失败: %s", err.Error())
	}

	// 监听配置变化
	// err = configClient.ListenConfig(vo.ConfigParam{
	// 	DataId: "user-web.json",
	// 	Group:  "dev",
	// 	OnChange: func(namespace, group, dataId, data string) {
	// 		fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
	// 	},
	// })

	// time.Sleep(3000 * time.Second)

}
