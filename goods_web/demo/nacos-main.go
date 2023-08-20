package main

import (
	"encoding/json"
	"fmt"
	"mxshop_api/goods_web/config"
	"time"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

// "github.com/nacos-group/nacos-sdk-go/"

// "github.com/nacos-group/nacos-sdk-go/common/constant"

func main() {
	sc := []constant.ServerConfig{
		{
			IpAddr: "go-nacos",
			Port:   8848,
		},
	}
	cc := constant.ClientConfig{
		NamespaceId:         "98a14936-c8c5-4da6-966e-25c75305326b",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
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
		DataId: "user-web.json",
		Group:  "dev"})

	if err != nil {
		panic(err)
	}

	fmt.Println(content)
	serverConfig := config.ServerConfig{}
	json.Unmarshal([]byte(content), &serverConfig)
	fmt.Println(serverConfig)

	// 监听配置变化
	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: "user-web.json",
		Group:  "dev",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
		},
	})

	time.Sleep(3000 * time.Second)

}
