package tiktoklib

import "github.com/easycb/easycb-go/tiktok"

// GetShopInfo TikTok获取店铺信息
func GetShopInfo(base *BaseTTService) *tiktok.GetActiveShopsRsp {
	// 从池中获取 Client
	client := base.GetClient()
	defer base.PutClient(client) // 确保归还
	client.SetAccessToken(base.shop.AccessToken)
	shops, err := client.GetActiveShops()
	if err != nil {
		return nil
	}
	return shops
}
