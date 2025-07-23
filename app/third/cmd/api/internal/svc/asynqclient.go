package svc

import "github.com/hibiken/asynq"

func NewAsynqClient(host, pass string, db int) *asynq.Client {
	return asynq.NewClient(
		asynq.RedisClientOpt{Addr: host, Password: pass, DB: db},
	)
}
