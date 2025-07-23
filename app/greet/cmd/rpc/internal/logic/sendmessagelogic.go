package logic

import (
	"context"
	"github.com/syyongx/php2go"
	"jekka-api-go/app/greet/cmd/rpc/greet"
	"jekka-api-go/app/greet/cmd/rpc/internal/svc"
	"jekka-api-go/pkg/services/util"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMessageLogic {
	return &SendMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendMessageLogic) SendMessage(in *greet.SendMessageReq) (*greet.SendMessageResp, error) {
	logx.WithContext(l.ctx).Infof("SendMessage: %+v", in)
	logx.WithContext(l.ctx).Infof("sonyflkeId: %+v", util.GenId())
	return &greet.SendMessageResp{
		Data: &greet.SendMessage{
			Status:  greet.Status_SUCCESS,
			Array:   []string{in.Message},
			Map:     map[string]int32{"timestamp": int32(php2go.Time())},
			Boolean: true,
		},
	}, nil
}
