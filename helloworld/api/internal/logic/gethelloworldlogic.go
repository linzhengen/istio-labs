package logic

import (
	"context"
	"helloworld/api/internal/svc"
	"helloworld/api/internal/types"
	"helloworld/rpc/helloworld"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHelloWorldLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetHelloWorldLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHelloWorldLogic {
	return &GetHelloWorldLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetHelloWorldLogic) GetHelloWorld(req *types.Request) (resp *types.Response, err error) {
	result, err := l.svcCtx.HelloWorldClient.Call(l.ctx, &helloworld.Request{})
	if err != nil {
		return nil, err
	}
	return &types.Response{Name: result.Name}, nil
}
