package logic

import (
	"context"
	"github.com/bytehello/gcc-zero/internal/bizerror"

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
	err := l.svcCtx.KvModel.Delete(in.Id)
	if err != nil {
		return nil, bizerror.Newf(bizerror.ErrCodeKvDelete, "KvDel err: %s", err.Error())
	}
	return &cc.KvDelReply{Pong: ""}, nil
}
