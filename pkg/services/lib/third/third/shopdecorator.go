package third

import (
	"encoding/json"
	"fmt"
	thirdconst "jekka-api-go/pkg/constant/third"
	"jekka-api-go/pkg/db/model"
)

type JkShopDecorator struct {
	Shop *model.JkShop
}

// ------------------------------------- third_session 相关 ---------------------------------------------

// SessionParser 定义接口
type SessionParser interface {
	GetAccessToken() string
	GetRefreshToken() string
	IsCrossBorder() bool
}

type SessionTT struct {
	Shop struct {
		SellerType string `json:"seller_type"`
	} `json:"shop"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type SessionSE struct {
	ERP struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	} `json:"erp"`
	Shop struct {
		IsCB   bool   `json:"is_cb"`
		Region string `json:"region"`
	} `json:"shop"`
}

type SessionLZD struct {
	Shop struct {
		CB    bool   `json:"cb"`
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"shop"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (t SessionTT) GetAccessToken() string   { return t.AccessToken }
func (t SessionTT) GetRefreshToken() string  { return t.RefreshToken }
func (t SessionTT) IsCrossBorder() bool      { return t.Shop.SellerType == "CROSS_BORDER" }
func (s SessionSE) GetAccessToken() string   { return s.ERP.AccessToken }
func (s SessionSE) GetRefreshToken() string  { return s.ERP.RefreshToken }
func (s SessionSE) IsCrossBorder() bool      { return s.Shop.IsCB }
func (l SessionLZD) GetAccessToken() string  { return l.AccessToken }
func (l SessionLZD) GetRefreshToken() string { return l.RefreshToken }
func (l SessionLZD) IsCrossBorder() bool     { return l.Shop.CB }

// ParseThirdSession 解析函数
func (d *JkShopDecorator) ParseThirdSession(thirdType int) (SessionParser, error) {
	switch thirdType {
	case thirdconst.ThirdTypeTikTok:
		var tt SessionTT
		if err := json.Unmarshal([]byte(d.Shop.ThirdSession), &tt); err == nil && tt.AccessToken != "" {
			return tt, nil
		}
	case thirdconst.ThirdTypeShopee:
		var se SessionSE
		if err := json.Unmarshal([]byte(d.Shop.ThirdSession), &se); err == nil && se.ERP.AccessToken != "" {
			return se, nil
		}
	case thirdconst.ThirdTypeLazada:
		var lzd SessionLZD
		if err := json.Unmarshal([]byte(d.Shop.ThirdSession), &lzd); err == nil && lzd.AccessToken != "" {
			return lzd, nil
		}
	default:
		return nil, fmt.Errorf("不支持的第三方类型")
	}
	return nil, fmt.Errorf("无法解析 ThirdSession")
}
