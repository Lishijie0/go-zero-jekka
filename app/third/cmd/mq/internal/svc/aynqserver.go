package svc

import (
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
	"jekka-api-go/pkg/constant"
)

func NewAsynqServer(host, pass string, db int) *asynq.Server {
	return asynq.NewServer(
		asynq.RedisClientOpt{Addr: host, Password: pass, DB: db},
		asynq.Config{
			IsFailure: func(err error) bool {
				logx.Debugf("[task-job] job exec failure, \n\terr: %+v", err)
				return true
			},
			Concurrency: 50,
			Queues: map[string]int{
				constant.QueueDefault: 5,
				constant.QueueLow:     15,
				constant.QueueHigh:    30,
			},
		},
	)
}
