package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"jekka-api-go/app/third/cmd/rpc/internal/config"
	"jekka-api-go/app/third/model"
)

type ServiceContext struct {
	Config config.Config

	ProductModel model.JkThirdProductModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:       c,
		ProductModel: model.NewJkThirdProductModel(sqlx.NewMysql(c.DB.DataSource)),
	}
}
