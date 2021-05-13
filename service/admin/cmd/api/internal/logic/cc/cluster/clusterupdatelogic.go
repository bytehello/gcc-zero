package logic

import (
	"context"

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
	// todo: add your logic here and delete this line

	return &types.ClusterUpdateReply{}, nil
}
