package logic

import (
	"context"

	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/cc"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type KvUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewKvUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KvUpdateLogic {
	return &KvUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *KvUpdateLogic) KvUpdate(in *cc.KvUpdateReq) (*cc.KvUpdateReply, error) {
	// todo: add your logic here and delete this line

	return &cc.KvUpdateReply{}, nil
}
