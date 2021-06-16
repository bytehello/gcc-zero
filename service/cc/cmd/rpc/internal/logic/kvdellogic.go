package logic

import (
	"context"

	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/cc"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type KvDelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewKvDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KvDelLogic {
	return &KvDelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *KvDelLogic) KvDel(in *cc.KvDelReq) (*cc.KvDelReply, error) {
	// todo: add your logic here and delete this line

	return &cc.KvDelReply{}, nil
}
