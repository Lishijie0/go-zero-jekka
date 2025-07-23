package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Redis struct {
	redis.RedisConf
	DB int `json:",default=0"`
}

type Config struct {
	rest.RestConf
	Redis Redis
}
