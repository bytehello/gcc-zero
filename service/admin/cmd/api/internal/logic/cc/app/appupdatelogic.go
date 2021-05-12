package logic

import (
	"context"
	"github.com/bytehello/gcc-zero/common/errorx"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/ccclient"
	"github.com/jinzhu/copier"

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
	var upReq ccclient.AppUpdateReq
	_ = copier.Copy(&upReq, req)
	_, err := l.svcCtx.CcRpcClient.AppUpdate(l.ctx, &upReq)
	if err != nil {
		return nil, errorx.DefaultCodeError(err.Error())
	}
	return &types.AppUpdateReply{
		Code:    "0",
		Message: "success",
	}, nil
}
