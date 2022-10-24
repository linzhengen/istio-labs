package logic

import (
	"context"
	"math/rand"
	"time"

	"helloworld/rpc/helloworld"
	"helloworld/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SayHelloLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSayHelloLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SayHelloLogic {
	return &SayHelloLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SayHelloLogic) SayHello(in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	l.Infof("Received: %v", in.GetName())
	rand.Seed(time.Now().UnixNano())
	min := 600
	max := 670
	s := rand.Intn(max-min+1) + min
	time.Sleep(time.Duration(s) * time.Millisecond)
	return &helloworld.HelloReply{Message: in.Name}, nil
}
