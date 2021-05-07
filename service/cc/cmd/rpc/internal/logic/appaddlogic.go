package logic

import (
	"context"

	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/cc"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type AppAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAppAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppAddLogic {
	return &AppAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AppAddLogic) AppAdd(in *cc.AppAddReq) (*cc.AppAddReply, error) {
	// todo: add your logic here and delete this line

	return &cc.AppAddReply{}, nil
}
