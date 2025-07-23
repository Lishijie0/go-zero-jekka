package main

import (
	"flag"
	"fmt"

	"jekka-api-go/app/third/cmd/rpc/internal/config"
	"jekka-api-go/app/third/cmd/rpc/internal/server"
	"jekka-api-go/app/third/cmd/rpc/internal/svc"
	"jekka-api-go/app/third/cmd/rpc/third"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/third.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		third.RegisterThirdServer(grpcServer, server.NewThirdServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
