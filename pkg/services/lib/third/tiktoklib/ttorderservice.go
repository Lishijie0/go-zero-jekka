package tiktoklib

import (
	"github.com/easycb/easycb-go"
	"github.com/easycb/easycb-go/tiktok"
)

// GetOrderInfo TikTok订单信息
func GetOrderInfo(base *BaseTTService, query easycb.AnyMap, config ApiQueryConfig) *tiktok.GetOrderDetailRsp {
	// 从池中获取 Client
	client := base.GetClient()
	defer base.PutClient(client) // 确保归还
	client.SetAccessToken(config.AccessToken)
	client.SetShopCipher(config.ShopCipher)
	orderDetail, err := client.GetOrderDetail(query)
	if err != nil {
		return nil
	}
	return orderDetail
}
