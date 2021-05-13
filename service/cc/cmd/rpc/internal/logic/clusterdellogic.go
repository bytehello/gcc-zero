package logic

import (
	"context"
	"github.com/bytehello/gcc-zero/common/errorx"

	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/cc"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type ClusterDelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClusterDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClusterDelLogic {
	return &ClusterDelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClusterDelLogic) ClusterDel(in *cc.ClusterDelReq) (*cc.ClusterDelReply, error) {
	if err := l.svcCtx.ClusterModel.Delete(in.Id); err != nil {
		return nil, errorx.DefaultCodeError(err.Error())
	}
	return &cc.ClusterDelReply{}, nil
}
