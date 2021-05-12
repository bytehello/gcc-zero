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
	var addReq ccclient.AppAddReq
	_ = copier.Copy(&addReq, req)
	_, err := l.svcCtx.CcRpcClient.AppAdd(l.ctx, &addReq)
	if err != nil {
		return nil, errorx.DefaultCodeError(err.Error())
	}
	return &types.AppAddReply{
		Code:    "0",
		Message: "success",
	}, nil
}
