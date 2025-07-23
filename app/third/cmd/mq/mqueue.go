package main

import (
	"context"
	"flag"
	"jekka-api-go/app/third/cmd/mq/internal/config"
	"jekka-api-go/app/third/cmd/mq/internal/handler"
	"jekka-api-go/app/third/cmd/mq/internal/monitor"
	"jekka-api-go/app/third/cmd/mq/internal/svc"
	"os"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/mqueue.yaml", "Specify the config file")

func main() {
	flag.Parse()
	var c config.Config

	conf.MustLoad(*configFile, &c, conf.UseEnv())

	// log、prometheus、trace、metricsUrl
	if err := c.SetUp(); err != nil {
		panic(err)
	}

	//logx.DisableStat()

	// 启动监控服务
	if err := monitor.StartMonitor(c); err != nil {
		logx.Errorf("Asynq monitor run err: %+v", err)
	}
	logx.Infof("Asynq monitor server start at 0.0.0.0:%d/monitor", c.Monitor.Port)

	// 注册job并启动
	svcContext, _ := svc.NewServiceContext(c)
	ctx := context.Background()
	cronJob := handler.NewCronJob(ctx, svcContext)
	mux := cronJob.Register()
	if err := svcContext.AsynqServer.Run(mux); err != nil {
		logx.WithContext(ctx).Errorf("!!!CronJobErr!!! run err:%+v", err)
		os.Exit(1)
	}
}
