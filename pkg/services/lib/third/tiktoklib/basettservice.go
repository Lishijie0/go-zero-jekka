package tiktoklib

import (
	"github.com/easycb/easycb-go/tiktok"
	"jekka-api-go/pkg/db/model"
	//third2 "jekka-api-go/pkg/services/third/third"
	"jekka-api-go/pkg/types/thirdtype"
	"sync"
)

// ApiQueryConfig 调用api的参数
type ApiQueryConfig struct {
	AccessToken string
	ShopCipher  string
}

type AppConfigTT struct {
	ClientID     string
	ClientSecret string
}

type BaseTTService struct {
	shop        model.JkShop
	clientPool  sync.Pool // Client 连接池
	httpOptions thirdtype.HttpOptions
}

type ThirdSessionTT struct {
	Shop struct {
		SellerType string `json:"seller_type"`
		ShopCipher string `json:"cipher"`
	} `json:"shop"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// NewTiktokClient 创建一个新的实例
func NewTiktokClient(appKey, appSecret string) *tiktok.Client {
	TiktokClient, _ := tiktok.NewClientDefault(appKey, appSecret)
	TiktokClient.SetBaseUrl("https://open-api.tiktokglobalshop.com")
	return TiktokClient
}

// NewBaseTTService 实例化
func NewBaseTTService(shop model.JkShop, config thirdtype.AppConfig) *BaseTTService {
	// 设置 httpOptions 的参数
	var httpOptions = thirdtype.HttpOptions{
		Timeout: 30,
		Verify:  false,
	}
	tiktok.WithTimeout(httpOptions.Timeout) // 设置超时时间

	//var thirdSession ThirdSessionTT
	//err := json.Unmarshal([]byte(shop.ThirdSession), &thirdSession)

	return &BaseTTService{
		shop: shop,
		clientPool: sync.Pool{
			New: func() interface{} {
				client := NewTiktokClient(config.ClientID, config.ClientSecret)
				//if err == nil && thirdSession.AccessToken != "" {
				//	client.SetAccessToken(thirdSession.AccessToken)
				//}
				return client
			},
		},
	}
}

// GetClient 从池中获取 Client
func (s *BaseTTService) GetClient() *tiktok.Client {
	return s.clientPool.Get().(*tiktok.Client)
}

// PutClient 归还 Client 到池中
func (s *BaseTTService) PutClient(client *tiktok.Client) {
	s.clientPool.Put(client)
}
