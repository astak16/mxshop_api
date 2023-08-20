package initialize

import (
	"fmt"
	"mxshop_api/goods_web/global"
	"mxshop_api/goods_web/proto"

	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func InitSrvConn() {
	consulInfo := global.ServerConfig.ConsulInfo
	conn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.UserSrvInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	)
	if err != nil {
		zap.S().Fatalf("[InitSrvConn] 连接 【商品服务失败】")
	}

	goodsSrvClient := proto.NewGoodsClient(conn)
	global.GoodsSrvClient = goodsSrvClient
}
