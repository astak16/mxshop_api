package initialize

import (
	"fmt"
	"mxshop_api/order_web/global"
	"mxshop_api/order_web/proto"

	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func InitSrvConn() {
	consulInfo := global.ServerConfig.ConsulInfo
	goodConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.GoodsSrvInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	)
	if err != nil {
		zap.S().Fatalf("[InitSrvConn] 连接 【商品服务失败】")
	}

	global.GoodsSrvClient = proto.NewGoodsClient(goodConn)

	orderConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.OrderSrvInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	)
	if err != nil {
		zap.S().Fatalf("[InitSrvConn] 连接 【订单服务失败】")
	}

	global.OrderSrvClient = proto.NewOrderClient(orderConn)

	invConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.InventorySrvInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	)
	if err != nil {
		zap.S().Fatalf("[InitSrvConn] 连接 【库存服务失败】")
	}

	global.InventorySrvClient = proto.NewInventoryClient(invConn)
}
