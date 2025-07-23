package cache

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"golang.org/x/net/context"
)

// BaseCache 用于处理缓存操作的基础结构体
type BaseCache struct {
	client *redis.Redis
	ctx    context.Context
}

// 定义缓存生命周期常量
const (
	TtlFiveSecond = 5
	TtlTenSecond  = 10
	TtlHalfMinute = 30
	TtlOneMinute  = 60
	TtlFiveMinute = 300
	TtlTenMinute  = 600
	TtlHalfHour   = 1800
	TtlOneHour    = 3600
	TtlTwoHour    = 7200
	TtlHalfDay    = 43200
	TtlOneDay     = 86400
	TtlThreeDay   = 259200
	TtlOneWeek    = 604800
	TtlOneMonth   = 2592000
)

// NewBaseCache 创建一个新的 BaseCache 实例
func NewBaseCache(redisAddr string, password string) *BaseCache {
	ctx := context.Background()
	conf := redis.RedisConf{
		Host: redisAddr,
		Pass: password,
		Type: "node", // 单节点模式
	}
	client := redis.MustNewRedis(conf)

	return &BaseCache{
		client: client,
		ctx:    ctx,
	}
}

// GetKey 返回格式化的缓存键
func (bc *BaseCache) GetKey(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}
