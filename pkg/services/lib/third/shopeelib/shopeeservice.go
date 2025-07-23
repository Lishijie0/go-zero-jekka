package shopeelib

import (
	"github.com/easycb/easycb-go/shopee"
)

// NewShopeeServer 创建一个新的 ShopeeServer 实例
func NewShopeeServer(appKey int64, appSecret string) *shopee.Client {
	ShopeeClient, _ := shopee.NewClientDefault(appKey, appSecret)
	return ShopeeClient
}
