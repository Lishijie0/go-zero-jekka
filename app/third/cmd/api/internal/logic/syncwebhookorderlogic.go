package logic

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"jekka-api-go/pkg/constant"
	"jekka-api-go/pkg/response/xerr"
	"time"

	"jekka-api-go/app/third/cmd/api/internal/svc"
	"jekka-api-go/app/third/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"jekka-api-go/app/third/cmd/mq/jobtype"
)

type SyncWebhookOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSyncWebhookOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SyncWebhookOrderLogic {
	return &SyncWebhookOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SyncWebhookOrderLogic) SyncWebhookOrder(req *types.SyncWebhookOrderReq) (resp *types.SyncWebhookOrderResp, err error) {
	// todo: add your logic here and delete this line
	payload, err := json.Marshal(jobtype.WebhookOrderPayload{JkShopId: req.JkShopId, JkUserId: req.JkUserId, OrderId: req.OrderId})
	if err != nil {
		return nil, xerr.NewErr(xerr.ServerCommonError, "解析payload失败，err:%v", err)
	}
	// 延迟队列
	_, err = l.svcCtx.AsynqClient.Enqueue(
		asynq.NewTask(jobtype.SyncWebhookOrder, payload),
		asynq.Queue(constant.QueueHigh), // 投递到队列 “high”
		asynq.ProcessIn(time.Second*1),  // 2s秒后执行
	)
	if err != nil {
		return nil, xerr.NewErr(xerr.ServerCommonError, "投递队列失败,err:%v", err)
	}
	return
}
