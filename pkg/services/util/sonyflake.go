package util

import (
	"github.com/sony/sonyflake"
	"github.com/zeromicro/go-zero/core/logx"
)

var flake *sonyflake.Sonyflake

func init() {
	flake = sonyflake.NewSonyflake(sonyflake.Settings{})
}

// GenId 相比snowflake(64bit)，sonyflake(63bit)1秒最多生成2.56w个，比Snowflake的409.6w少太多
// 如果感觉不够用，目前的解决方案是跑多个实例生成同一业务的ID来弥补
func GenId() int64 {

	id, err := flake.NextID()
	if err != nil {
		logx.Severef("flake NextID failed with %s \n", err)
		panic(err)
	}

	return int64(id)
}
