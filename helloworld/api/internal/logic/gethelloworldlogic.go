package logic

import (
	"context"

	"helloworld/api/internal/svc"
	"helloworld/api/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
