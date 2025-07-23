package logic

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	jobtype2 "jekka-api-go/app/third/cmd/mq/jobtype"
	"jekka-api-go/pkg/constant"
	"jekka-api-go/pkg/response/xerr"
	"time"

	"jekka-api-go/app/third/cmd/api/internal/svc"
	"jekka-api-go/app/third/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AsyncShopBasicDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAsyncShopBasicDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AsyncShopBasicDataLogic {
	return &AsyncShopBasicDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AsyncShopBasicDataLogic) AsyncShopBasicData(req *types.AsyncShopBasicDataReq) (resp *types.AsyncShopBasicDataResp, err error) {
	payload, err := json.Marshal(jobtype2.SyncShopBasicPayload{JkShopId: req.JkShopId, JkUserId: req.JkUserId})
	if err != nil {
		return nil, xerr.NewErr(xerr.ServerCommonError, "解析payload失败，err:%v", err)
	}
	//_, err = l.svcCtx.AsynqClient.Enqueue(asynq.NewTask(jobtype.SyncShopBasicData, payload), asynq.Queue(xconstant.QueueHigh))
	// 延迟队列
	_, err = l.svcCtx.AsynqClient.Enqueue(
		asynq.NewTask(jobtype2.SyncShopBasicData, payload),
		asynq.Queue(constant.QueueHigh), // 投递到队列 “high”
		asynq.ProcessIn(time.Second*2),  // 2s秒后执行
	)
	if err != nil {
		return nil, xerr.NewErr(xerr.ServerCommonError, "投递队列失败,err:%v", err)
	}
	return &types.AsyncShopBasicDataResp{}, nil
}
