package tools

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type FreqUtil struct {
	RedisClient *redis.Redis
	Prefix      string
}

// NewFreqUtil 创建一个默认配置的 FreqUtil 实例，隐藏 Redis 初始化逻辑
func NewFreqUtil() *FreqUtil {
	host := os.Getenv("REDIS_HOST")
	pass := os.Getenv("REDIS_PASSWORD")
	// 内部封装 Redis 配置
	conf := redis.RedisConf{
		Host: host,
		Type: "node",
		Pass: pass,
	}

	rds, err := redis.NewRedis(conf)
	if err != nil {
		return nil
	}

	return &FreqUtil{
		RedisClient: rds,
		Prefix:      "freq_util:",
	}
}

// Control 函数用于记录 message 频率，返回当前累计的 count。
// message: 要记录的 key，比如错误消息
// ex: 过期时间（单位：秒）
func (f *FreqUtil) Control(ctx context.Context, message string, ex int) (int, error) {
	key := f.Prefix + message

	// Redis INCR 操作
	count, err := f.RedisClient.IncrCtx(ctx, key)
	if err != nil {
		return 0, fmt.Errorf("redis INCR err: %v", err)
	}

	if count == 1 {
		// 第一次写入，设置过期时间
		_ = f.RedisClient.ExpireCtx(ctx, key, int(time.Duration(ex)*time.Second))
	}

	return int(count), nil
}
