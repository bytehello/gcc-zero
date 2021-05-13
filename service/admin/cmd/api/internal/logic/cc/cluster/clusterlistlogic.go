package logic

import (
	"context"

	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/svc"
	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ClusterListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClusterListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ClusterListLogic {
	return ClusterListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClusterListLogic) ClusterList(req types.ClusterListReq) (*types.ClusterListReply, error) {
	// todo: add your logic here and delete this line

	return &types.ClusterListReply{}, nil
}
