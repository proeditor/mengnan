package main

import (
	"flag"
	"fmt"

	"github.com/proeditor/mengnan/college/rpc/college-rpc"
	"github.com/proeditor/mengnan/college/rpc/internal/config"
	"github.com/proeditor/mengnan/college/rpc/internal/server"
	"github.com/proeditor/mengnan/college/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/college.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		college_rpc.RegisterCollegeServer(grpcServer, server.NewCollegeServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
