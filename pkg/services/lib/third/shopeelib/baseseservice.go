package shopeelib

import (
	"encoding/json"
	"github.com/easycb/easycb-go/shopee"
	"jekka-api-go/pkg/db/model"
	"jekka-api-go/pkg/types/thirdtype"
	"sync"
)

type AppConfigSE struct {
	ClientID     int64
	ClientSecret string
}

type BaseSEService struct {
	shop        model.JkShop
	clientPool  sync.Pool // Client 连接池
	httpOptions thirdtype.HttpOptions
}

type ThirdSessionSE struct {
	ERP struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	} `json:"erp"`
	Shop struct {
		IsCB   bool   `json:"is_cb"`
		Region string `json:"region"`
	} `json:"shop"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// NewShopeeClient 创建一个新的实例
func NewShopeeClient(appKey int64, appSecret string) *shopee.Client {
	ShopeeClient, _ := shopee.NewClientDefault(appKey, appSecret)
	return ShopeeClient
}

// NewBaseSEService 实例化
func NewBaseSEService(shop model.JkShop, config AppConfigSE) *BaseSEService {
	// 设置 httpOptions 的参数
	var httpOptions = thirdtype.HttpOptions{
		Timeout: 30,
		Verify:  false,
	}
	shopee.WithTimeout(httpOptions.Timeout) // 设置超时时间

	var thirdSession ThirdSessionSE
	err := json.Unmarshal([]byte(shop.ThirdSession), &thirdSession)

	return &BaseSEService{
		shop: shop,
		clientPool: sync.Pool{
			New: func() interface{} {
				client := NewShopeeClient(config.ClientID, config.ClientSecret)
				if err == nil && thirdSession.AccessToken != "" {
					client.SetAccessToken(thirdSession.AccessToken)
				}
				return client
			},
		},
	}
}

// GetClient 从池中获取 Client
func (s *BaseSEService) GetClient() *shopee.Client {
	return s.clientPool.Get().(*shopee.Client)
}

// PutClient 归还 Client 到池中
func (s *BaseSEService) PutClient(client *shopee.Client) {
	s.clientPool.Put(client)
}
