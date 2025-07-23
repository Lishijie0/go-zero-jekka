package config

import (
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

// Redis 自定义Redis配置，添加DB选择
type Redis struct {
	redis.RedisConf
	DB int `json:",default=0"`
}

// Monitor 监控配置
type Monitor struct {
	Enabled bool   `json:",default=true"`
	Port    int    `json:",default=8080"`
	Path    string `json:",default=/monitor"`
	Name    string
	Pass    string
}

type MySQL struct {
	DataSource      string
	MaxIdleConns    int `json:",default=10"`
	MaxOpenConns    int `json:",default=100"`
	ConnMaxLifetime int `json:",default=3600"`
}

type MySQLThird struct {
	DataSource      string
	MaxIdleConns    int `json:",default=10"`
	MaxOpenConns    int `json:",default=100"`
	ConnMaxLifetime int `json:",default=3600"`
}

type MySQLThirdV2 struct {
	DataSource      string
	MaxIdleConns    int `json:",default=10"`
	MaxOpenConns    int `json:",default=100"`
	ConnMaxLifetime int `json:",default=3600"`
}

// ThirdPlatformConfig 第三方平台配置
type ThirdPlatformConfig struct {
	Lazada    LazadaConfig    `json:"lazada"`
	Shopee    ShopeeConfig    `json:"shopee"`
	ShopeeErp ShopeeErpConfig `json:"shopeeErp"`
	Tiktok    TiktokConfig    `json:"tiktok"`
	TiktokUs  TiktokUsConfig  `json:"tiktokUs"`
}

// LazadaConfig 用于存储 Lazada 配置
type LazadaConfig struct {
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
}

// ShopeeConfig 用于存储 Shopee 配置
type ShopeeConfig struct {
	ClientID     int64  `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
}

// ShopeeErpConfig 用于存储 Shopee 配置
type ShopeeErpConfig struct {
	ClientID     int64  `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
}

// TiktokConfig 用于存储 Lazada 配置
type TiktokConfig struct {
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
}

// TiktokUsConfig 用于存储 Lazada 配置
type TiktokUsConfig struct {
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
}

type Config struct {
	service.ServiceConf

	Redis Redis

	Monitor Monitor

	MySQL MySQL

	MySQLThird MySQLThird

	MySQLThirdV2 MySQLThirdV2

	ThirdPlatformConfig ThirdPlatformConfig
}
