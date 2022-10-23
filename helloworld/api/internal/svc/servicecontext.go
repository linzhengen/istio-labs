package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"helloworld/api/internal/config"
	"helloworld/rpc/greeter"
	"time"
)

type ServiceContext struct {
	Config                  config.Config
	HelloWorldGreeterClient greeter.Greeter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		HelloWorldGreeterClient: greeter.NewGreeter(
			zrpc.MustNewClient(zrpc.NewDirectClientConf([]string{c.HellWorldEndpoint}, "", ""), zrpc.WithTimeout(10*time.Second)),
		),
	}
}
