package logic

import (
	"context"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/cc"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type KvGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewKvGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KvGetLogic {
	return &KvGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *KvGetLogic) KvGet(in *cc.KvGetReq) (*cc.KvGetReq, error) {
	return &cc.KvGetReq{}, nil
}
