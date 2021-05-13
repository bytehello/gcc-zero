package logic

import (
	"context"

	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/svc"
	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ClusterDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClusterDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) ClusterDelLogic {
	return ClusterDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClusterDelLogic) ClusterDel(req types.ClusterDelReq) (*types.ClusterDelReply, error) {
	// todo: add your logic here and delete this line

	return &types.ClusterDelReply{}, nil
}
