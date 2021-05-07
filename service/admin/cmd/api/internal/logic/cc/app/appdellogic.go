package logic

import (
	"context"

	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/svc"
	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AppDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAppDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) AppDelLogic {
	return AppDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AppDelLogic) AppDel(req types.AppDelReq) (*types.AppDelReply, error) {
	// todo: add your logic here and delete this line

	return &types.AppDelReply{}, nil
}
