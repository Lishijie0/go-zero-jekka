package lazada

import (
	"encoding/json"
	"github.com/easycb/easycb-go/lazada"
	"jekka-api-go/pkg/db/model"
	"jekka-api-go/pkg/types/thirdtype"
	"sync"
)

type AppConfigLZD struct {
	ClientID     string
	ClientSecret string
}

type BaseLZDService struct {
	shop        model.JkShop
	clientPool  sync.Pool // Client 连接池
	httpOptions thirdtype.HttpOptions
}

type ThirdSessionLZD struct {
	Shop struct {
		IsCB  bool   `json:"cb"`
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"shop"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// NewLazadaClient 创建一个新的实例
func NewLazadaClient(appKey, appSecret string) *lazada.Client {
	LazadaClient, _ := lazada.NewClientDefault(appKey, appSecret)
	LazadaClient.SetBaseUrl("https://open-api.tiktokglobalshop.com")
	return LazadaClient
}

// NewBaseLZDService 实例化
func NewBaseLZDService(shop model.JkShop, config AppConfigLZD) *BaseLZDService {
	// 设置 httpOptions 的参数
	var httpOptions = thirdtype.HttpOptions{
		Timeout: 30,
		Verify:  false,
	}
	lazada.WithTimeout(httpOptions.Timeout) // 设置超时时间

	var thirdSession ThirdSessionLZD
	err := json.Unmarshal([]byte(shop.ThirdSession), &thirdSession)

	return &BaseLZDService{
		shop: shop,
		clientPool: sync.Pool{
			New: func() interface{} {
				client := NewLazadaClient(config.ClientID, config.ClientSecret)
				if err == nil && thirdSession.AccessToken != "" {
					client.SetAccessToken(thirdSession.AccessToken)
				}
				return client
			},
		},
	}
}

// GetClient 从池中获取 Client
func (s *BaseLZDService) GetClient() *lazada.Client {
	return s.clientPool.Get().(*lazada.Client)
}

// PutClient 归还 Client 到池中
func (s *BaseLZDService) PutClient(client *lazada.Client) {
	s.clientPool.Put(client)
}
