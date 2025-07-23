package handler

import (
	"context"
	"github.com/hibiken/asynq"
	"jekka-api-go/app/third/cmd/mq/internal/svc"
	"jekka-api-go/app/third/cmd/mq/jobtype"
)

type CronJob struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCronJob(ctx context.Context, svcCtx *svc.ServiceContext) *CronJob {
	return &CronJob{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Register register job
func (l *CronJob) Register() *asynq.ServeMux {
	mux := asynq.NewServeMux()
	mux.Handle(jobtype.SyncShopBasicData, NewSyncShopOrderHandler(l.svcCtx))
	mux.Handle(jobtype.SyncWebhookOrder, NewSyncWebhookOrderHandler(l.svcCtx))
	return mux
}
