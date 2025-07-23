package config

import (
	"fmt"
	"jekka-api-go/pkg/services/util"
	"sync"

	"github.com/zeromicro/go-zero/core/conf"
)

// AppConfig 定义整个配置结构，对应 config.yaml
type AppConfig struct {
	Name          string              `yaml:"Name"`
	Mode          string              `yaml:"Mode"`
	Log           map[string]any      `yaml:"Log"`
	Redis         RedisConfig         `yaml:"Redis"`
	MySQL         MySQLConfig         `yaml:"MySQL"`
	MySQLThird    MySQLConfig         `yaml:"MySQLThird"`
	MySQLThirdV2  MySQLConfig         `yaml:"MySQLThirdV2"`
	Monitor       MonitorConfig       `yaml:"Monitor"`
	ThirdPlatform ThirdPlatformConfig `yaml:"ThirdPlatformConfig"`
}

type RedisConfig struct {
	Host string `yaml:"Host"`
	Type string `yaml:"Type"`
	Pass string `yaml:"Pass"`
	DB   int    `yaml:"DB"`
}

type MySQLConfig struct {
	DataSource      string `yaml:"DataSource"`
	MaxIdleConns    int    `yaml:"MaxIdleConns"`
	MaxOpenConns    int    `yaml:"MaxOpenConns"`
	ConnMaxLifetime int    `yaml:"ConnMaxLifetime"`
}

type MonitorConfig struct {
	Enabled bool   `yaml:"Enabled"`
	Port    int    `yaml:"Port"`
	Path    string `yaml:"Path"`
	Name    string `yaml:"Name"`
	Pass    string `yaml:"Pass"`
}

type ThirdPlatformConfig struct {
	Lazada    map[string]string `yaml:"Lazada"`
	Shopee    map[string]string `yaml:"Shopee"`
	ShopeeErp map[string]string `yaml:"ShopeeErp"`
	TikTok    map[string]string `yaml:"TikTok"`
	TikTokUs  map[string]string `yaml:"TikTokUs"`
}

// 全局配置单例
var (
	globalConfig *AppConfig
	once         sync.Once
)

// GetGlobalConfig 获取全局配置（只加载一次）
func GetGlobalConfig(path string) *AppConfig {
	once.Do(func() {
		var err error
		globalConfig, err = LoadConfigYaml(path)
		if err != nil {
			panic("加载配置失败: " + err.Error())
		}
	})
	return globalConfig
}

// LoadConfigYaml 加载 YAML 配置文件
func LoadConfigYaml(path string) (*AppConfig, error) {
	cfg := new(AppConfig)

	if err := conf.Load(path, cfg); err != nil {
		return nil, fmt.Errorf("加载 config.yaml 失败: %v", err)
	}

	return cfg, nil
}

// NewConfig 单例构造函数，只会加载一次 config.yaml
func NewConfig() *AppConfig {
	path := "../config.yaml"
	// 检测文件是否存在
	if !util.FileExists(path) {
		panic("config.yaml 文件不存在")
	}
	once.Do(func() {
		cfg := new(AppConfig)
		if err := conf.Load(path, cfg); err != nil {
			panic(fmt.Sprintf("加载 config.yaml 失败: %v", err))
		}
		globalConfig = cfg
	})
	return globalConfig
}

//// RedisClient 全局 Redis 实例（可选）
//var (
//	redisClient *redis.Redis
//	redisOnce   sync.Once
//)
//
//func GetGlobalRedis() *redis.Redis {
//	redisOnce.Do(func() {
//		cfg := GetGlobalConfig("../../config.yaml")
//		redisConf := redis.RedisConf{
//			Host: cfg.Redis.Host,
//			Type: cfg.Redis.Type,
//			Pass: cfg.Redis.Pass,
//		}
//		client, err := redis.NewRedis(redisConf)
//		if err != nil {
//			panic("初始化 Redis 失败: " + err.Error())
//		}
//		redisClient = client
//	})
//	return redisClient
//}
