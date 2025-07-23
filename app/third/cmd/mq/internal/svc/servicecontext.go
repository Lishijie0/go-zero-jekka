package svc

import (
	"fmt"
	"jekka-api-go/app/third/cmd/mq/internal/config"

	"github.com/hibiken/asynq"
	"gorm.io/gorm"
	"gorm.io/sharding"
)

type ServiceContext struct {
	Config      config.Config
	AsynqServer *asynq.Server
	DB          *gorm.DB
	DBThird     *gorm.DB
	DBThirdV2   *gorm.DB
}

func NewServiceContext(c config.Config) (*ServiceContext, error) {
	db, err := NewDB(c.MySQL)
	if err != nil {
		return nil, err
	}

	dbThird, err := NewDBThird(c.MySQLThird)
	if err != nil {
		return nil, err
	}

	dbThirdV2, err := NewDBThirdV2(c.MySQLThirdV2)
	if err != nil {
		return nil, err
	}

	//dbThirdV2, err := gorm.Open(mysql.Open(c.MySQLThirdV2.DataSource), &gorm.Config{})
	//if err != nil {
	//	return nil, err
	//}
	// 注册分片插件
	err = dbThirdV2.Use(sharding.Register(sharding.Config{
		ShardingKey:         "jk_user_id",
		NumberOfShards:      64,
		PrimaryKeyGenerator: sharding.PKSnowflake,
		// 如果需要自定义分片规则，可以设置 ShardingAlgorithm 函数
		ShardingAlgorithm: func(val any) (suffix string, err error) {
			if uid, ok := val.(int64); ok {
				return fmt.Sprintf("_%02d", uid%64), nil
			}
			return "", fmt.Errorf("invalid jk_user_id")
		},
	}, "jk_third_order", "jk_third_order_detail", "jk_third_order_extra"))
	if err != nil {
		return nil, err
	}

	asynqServer := NewAsynqServer(c.Redis.Host, c.Redis.Pass, c.Redis.DB)

	return &ServiceContext{
		Config:      c,
		AsynqServer: asynqServer,
		DB:          db,
		DBThird:     dbThird,
		DBThirdV2:   dbThirdV2,
	}, nil

}
