package logic

import (
	"context"

	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/cc"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type ClusterUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClusterUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClusterUpdateLogic {
	return &ClusterUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClusterUpdateLogic) ClusterUpdate(in *cc.ClusterUpdateReq) (*cc.ClusterUpdateReply, error) {
	// todo: add your logic here and delete this line

	return &cc.ClusterUpdateReply{}, nil
}
