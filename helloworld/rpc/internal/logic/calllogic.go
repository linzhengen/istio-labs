package logic

import (
	"context"
	"math/rand"
	"time"

	"helloworld/rpc/helloworld"
	"helloworld/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CallLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCallLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CallLogic {
	return &CallLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CallLogic) Call(in *helloworld.Request) (*helloworld.Response, error) {
	rand.Seed(time.Now().UnixNano())
	min := 100
	max := 500
	s := rand.Intn(max-min+1) + min
	time.Sleep(time.Duration(s) * time.Millisecond)
	return &helloworld.Response{Name: "foo"}, nil
}
