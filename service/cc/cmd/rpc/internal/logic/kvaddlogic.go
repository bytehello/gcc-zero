package logic

import (
	"context"

	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/cc"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type KvAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewKvAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KvAddLogic {
	return &KvAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *KvAddLogic) KvAdd(in *cc.KvAddReq) (*cc.KvAddReply, error) {
	// todo: add your logic here and delete this line

	return &cc.KvAddReply{}, nil
}
