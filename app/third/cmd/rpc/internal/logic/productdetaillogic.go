package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"jekka-api-go/app/third/cmd/rpc/internal/svc"
	"jekka-api-go/app/third/cmd/rpc/third"
	"jekka-api-go/app/third/model"
	"jekka-api-go/pkg/response/xerr"
)

type ProductDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductDetailLogic {
	return &ProductDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ProductDetail 查询商品详情
func (l *ProductDetailLogic) ProductDetail(in *third.ProductDetailReq) (*third.ProductDetailResp, error) {
	logx.WithContext(l.ctx).Infof("ProductDetail: %+v", in)
	// 查询
	productDetail, err := l.svcCtx.ProductModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return &third.ProductDetailResp{
				Data: nil,
			}, nil
		}
		return nil, xerr.NewErr(xerr.DbError, "ProductDetail FindOneById db err: %v, in:%+v", err, in)
	}
	// 返回
	return &third.ProductDetailResp{
		Data: &third.ProductDetail{
			Id:        productDetail.Id,
			JkShopId:  productDetail.JkShopId,
			ProductId: productDetail.ProductId,
			Title:     productDetail.Title,
			MinPrice:  float32(productDetail.MinPrice),
			MaxPrice:  float32(productDetail.MaxPrice),
			CreateAt:  productDetail.CreatedAt.Unix(),
		},
	}, nil
}
