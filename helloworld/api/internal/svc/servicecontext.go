package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"helloworld/api/internal/config"
	"helloworld/rpc/helloworldclient"
)

type ServiceContext struct {
	Config           config.Config
	HelloWorldClient helloworldclient.HelloWorld
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		HelloWorldClient: helloworldclient.NewHelloWorld(
			zrpc.MustNewClient(zrpc.NewDirectClientConf([]string{c.HellWorldEndpoint}, "", "")),
		),
	}
}
