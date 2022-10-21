package logic

import (
	"context"

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
	return &helloworld.Response{Name: "foo"}, nil
}
