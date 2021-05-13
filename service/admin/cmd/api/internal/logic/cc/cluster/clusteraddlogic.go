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

type ClusterAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClusterAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) ClusterAddLogic {
	return ClusterAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClusterAddLogic) ClusterAdd(req types.ClusterAddReq) (*types.ClusterAddReply, error) {
	var addReq ccclient.ClusterAddReq
	_ = copier.Copy(&addReq, &req)
	reply, err := l.svcCtx.CcRpcClient.ClusterAdd(l.ctx, &addReq)
	if err != nil {
		return nil, errorx.DefaultCodeError(err.Error())
	}
	return &types.ClusterAddReply{
		Id:      reply.Id,
		Message: "success",
		Code:    "0",
	}, nil
}
