package logic

import (
	"context"

	"jekka-api-go/app/third/cmd/rpc/internal/svc"
	"jekka-api-go/app/third/cmd/rpc/third"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductListLogic {
	return &ProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductList 店铺商品列表
func (l *ProductListLogic) ProductList(in *third.ProductListReq) (*third.ProductListResp, error) {
	// 获取商品列表
	list, err := l.svcCtx.ProductModel.GetListByJkShopId(l.ctx, in.JkShopId, in.Page, in.PageSize)
	if err != nil {
		return nil, err
	}

	// 转换为响应格式
	resp := &third.ProductListResp{
		Data: make([]*third.ProductDetail, 0, len(list)),
	}
	logx.WithContext(l.ctx).Infof("遍历结果开始: %+v", list)
	// 遍历结果并转换
	for _, item := range list {
		resp.Data = append(resp.Data, &third.ProductDetail{
			Id:        item.Id,
			JkShopId:  item.JkShopId,
			ProductId: item.ProductId,
			Title:     item.Title,
			MinPrice:  float32(item.MinPrice),
			MaxPrice:  float32(item.MaxPrice),
			CreateAt:  item.CreatedAt.Unix(),
		})
	}

	return resp, nil
}
