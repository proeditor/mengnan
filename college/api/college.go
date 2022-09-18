package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest"

	"college/api/internal/config"
	"college/api/internal/handler"
	"college/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/college-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("开始启动 college api 服务 at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
