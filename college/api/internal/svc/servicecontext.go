package svc

import (
	"college/api/internal/config"
	"github.com/proeditor/mengnan/college/rpc/college"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		Colleger: college.NewCollege(zrpc.MustNewClient(c.College)), // 手动代码
	}
}
