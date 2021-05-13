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

type ClusterUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClusterUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) ClusterUpdateLogic {
	return ClusterUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClusterUpdateLogic) ClusterUpdate(req types.ClusterUpdateReq) (*types.ClusterUpdateReply, error) {
	var upReq ccclient.ClusterUpdateReq
	_ = copier.Copy(&upReq, &req)
	_, err := l.svcCtx.CcRpcClient.ClusterUpdate(l.ctx, &upReq)
	if err != nil {
		return nil, errorx.DefaultCodeError(err.Error())
	}
	return &types.ClusterUpdateReply{
		Code:    "0",
		Message: "success",
	}, nil
}
