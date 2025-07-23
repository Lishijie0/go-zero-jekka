package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
	"time"
)

var _ JkThirdProductModel = (*customJkThirdProductModel)(nil)

type (
	// JkThirdProductModel is an interface to be customized, add more methods here,
	// and implement the added methods in customJkThirdProductModel.
	JkThirdProductModel interface {
		jkThirdProductModel
		GetListByJkShopId(ctx context.Context, jkShopId int64, page int64, pageSize int64) ([]*JkThirdProduct, error)
	}

	customJkThirdProductModel struct {
		*defaultJkThirdProductModel
	}

	// ListJkThirdProduct 用于列表查询的精简结构体
	ListJkThirdProduct struct {
		Id        int64     `db:"id"`
		JkShopId  int64     `db:"jk_shop_id"`
		ProductId string    `db:"product_id"`
		Title     string    `db:"title"`
		MinPrice  float64   `db:"min_price"`
		MaxPrice  float64   `db:"max_price"`
		CreatedAt time.Time `db:"created_at"`
	}
)

// NewJkThirdProductModel returns a model for the database table.
func NewJkThirdProductModel(conn sqlx.SqlConn) JkThirdProductModel {
	return &customJkThirdProductModel{
		defaultJkThirdProductModel: newJkThirdProductModel(conn),
	}
}

// GetListByJkShopId 根据jekka店铺id获取商品列表
func (m *customJkThirdProductModel) GetListByJkShopId(ctx context.Context, jkShopId int64, page int64, pageSize int64) ([]*JkThirdProduct, error) {
	logx.WithContext(ctx).Infof("开始查询: %v", jkShopId)
	// 只查询需要的字段
	fields := []string{
		"`id`",
		"`jk_shop_id`",
		"`product_id`",
		"`title`",
		"`min_price`",
		"`max_price`",
		"`created_at`",
	}

	// 计算偏移量
	offset := (page - 1) * pageSize

	// 构建查询语句
	query := fmt.Sprintf("select %s from %s where `jk_shop_id` = ? limit ?,?",
		strings.Join(fields, ","),
		m.table,
	)

	// 定义结果切片，使用精简的结构体
	var list []*ListJkThirdProduct
	err := m.conn.QueryRowsCtx(ctx, &list, query, jkShopId, offset, pageSize)
	if err != nil {
		fmt.Printf("查询错误: %v\n", err)
		return nil, err
	}

	// 转换为完整的JkThirdProduct结构体
	resp := make([]*JkThirdProduct, 0, len(list))
	for _, item := range list {
		resp = append(resp, &JkThirdProduct{
			Id:        item.Id,
			JkShopId:  item.JkShopId,
			ProductId: item.ProductId,
			Title:     item.Title,
			MinPrice:  item.MinPrice,
			MaxPrice:  item.MaxPrice,
			CreatedAt: item.CreatedAt,
		})
	}

	return resp, nil
}
