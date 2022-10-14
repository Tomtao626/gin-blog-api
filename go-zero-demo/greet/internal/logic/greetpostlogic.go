package logic

import (
	"context"

	"go-zero-demo/greet/internal/svc"
	"go-zero-demo/greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GreetPostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGreetPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GreetPostLogic {
	return &GreetPostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GreetPostLogic) GreetPost(req *types.RequestJson) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	return &types.Response{
		Message: "Hello Go-Zero!" + req.Name,
	}, nil
}
