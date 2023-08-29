package shop_cart

import (
	"context"
	"fmt"
	"mxshop_api/order_web/api"
	"mxshop_api/order_web/forms"
	"mxshop_api/order_web/global"
	"mxshop_api/order_web/proto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func New(ctx *gin.Context) {
	itemForm := forms.ShopCartItemForm{}
	if err := ctx.ShouldBindJSON(&itemForm); err != nil {
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}
	_, err := global.GoodsSrvClient.GetGoodsDetail(context.Background(), &proto.GoodInfoRequest{
		Id: itemForm.GoodsId,
	})

	if err != nil {
		zap.S().Errorw("[New] 查询【商品信息】失败")
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}

	invRsp, err := global.InventorySrvClient.InvDetail(context.Background(), &proto.GoodsInvInfo{
		GoodsId: itemForm.GoodsId,
	})
	if err != nil {
		zap.S().Errorw("[New] 查询【库存信息】失败")
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}

	if invRsp.Nums < itemForm.Nums {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "库存不足",
		})
		return
	}

	userId, _ := ctx.Get("userId")
	rsp, err := global.OrderSrvClient.CreateCartItem(context.Background(), &proto.CartItemRequest{
		UserId:  int32(userId.(uint)),
		GoodsId: itemForm.GoodsId,
		Nums:    itemForm.Nums,
	})
	if err != nil {
		zap.S().Errorw("[New] 添加【添加购物车】失败")
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}

	reMap := map[string]interface{}{}
	reMap["id"] = rsp.Id
	ctx.JSON(http.StatusOK, reMap)
}

func List(ctx *gin.Context) {
	userId, _ := ctx.Get("userId")
	rsp, err := global.OrderSrvClient.CartItemList(ctx, &proto.UserInfo{
		Id: int32(userId.(uint)),
	})
	if err != nil {
		zap.S().Errorw("[List] 查询【购物车列表】失败")
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ids := make([]int32, 0)
	for _, item := range rsp.Data {
		ids = append(ids, item.GoodsId)
	}
	if len(ids) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"total": 0,
		})
		return
	}

	goodsRes, err := global.GoodsSrvClient.BatchGetGoods(context.Background(), &proto.BatchGoodsIdInfo{
		Id: ids,
	})
	if err != nil {
		zap.S().Errorw("[List] 批量查询【商品列表】失败")
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}

	reMap := gin.H{
		"total": rsp.Total,
	}

	goodList := make([]interface{}, 0)
	for _, item := range rsp.Data {
		for _, goods := range goodsRes.Data {
			if item.GoodsId == goods.Id {
				tmpMap := map[string]interface{}{}
				tmpMap["id"] = item.Id
				tmpMap["goods_id"] = item.GoodsId
				tmpMap["good_name"] = goods.Name
				tmpMap["good_image"] = goods.GoodsFrontImage
				tmpMap["good_price"] = goods.ShopPrice
				tmpMap["nums"] = item.Nums
				tmpMap["checked"] = item.Checked
				goodList = append(goodList, tmpMap)
			}
		}
	}
	reMap["data"] = goodList
	ctx.JSON(http.StatusOK, reMap)
}

func Update(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg": "url格式出错",
		})
		return
	}
	itemForm := forms.ShopCartItemUpdateForm{}
	if err := ctx.ShouldBindJSON(&itemForm); err != nil {
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}
	fmt.Println(itemForm)
	userId, _ := ctx.Get("userId")
	request := proto.CartItemRequest{
		UserId:  int32(userId.(uint)),
		GoodsId: int32(i),
		Nums:    itemForm.Nums,
		Checked: false,
	}
	if itemForm.Checked != nil {
		request.Checked = *itemForm.Checked
	}

	_, err = global.OrderSrvClient.UpdateCartItem(context.Background(), &request)

	if err != nil {
		zap.S().Errorw("[Update] 更新【购物车记录】失败")
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.Status(http.StatusOK)
}

func Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg": "url格式出错",
		})
		return
	}
	userId, _ := ctx.Get("userId")
	_, err = global.OrderSrvClient.DeleteCartItem(context.Background(), &proto.CartItemRequest{
		UserId:  int32(userId.(uint)),
		GoodsId: int32(i),
	})
	if err != nil {
		zap.S().Errorw("[Delete] 删除【购物车记录】失败")
		api.HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.Status(http.StatusOK)
}
