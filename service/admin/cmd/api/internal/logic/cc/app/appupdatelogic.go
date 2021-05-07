package logic

import (
	"context"

	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/svc"
	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AppUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAppUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) AppUpdateLogic {
	return AppUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AppUpdateLogic) AppUpdate(req types.AppUpdateReq) (*types.AppUpdateReply, error) {
	// todo: add your logic here and delete this line

	return &types.AppUpdateReply{}, nil
}
