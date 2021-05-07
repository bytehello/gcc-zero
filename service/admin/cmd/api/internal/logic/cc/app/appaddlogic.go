package logic

import (
	"context"

	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/svc"
	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AppAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAppAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) AppAddLogic {
	return AppAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AppAddLogic) AppAdd(req types.AppAddReq) (*types.AppAddReply, error) {
	// todo: add your logic here and delete this line

	return &types.AppAddReply{}, nil
}
