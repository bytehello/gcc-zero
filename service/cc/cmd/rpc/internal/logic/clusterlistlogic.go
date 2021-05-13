package logic

import (
	"context"

	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/cc"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type ClusterListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClusterListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClusterListLogic {
	return &ClusterListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClusterListLogic) ClusterList(in *cc.ClusterListReq) (*cc.ClusterListReply, error) {
	// todo: add your logic here and delete this line

	return &cc.ClusterListReply{}, nil
}
