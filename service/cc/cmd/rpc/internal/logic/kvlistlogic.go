package logic

import (
	"context"

	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/cc"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type KvListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewKvListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KvListLogic {
	return &KvListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *KvListLogic) KvList(in *cc.KvListReq) (*cc.KvListReply, error) {
	// todo: add your logic here and delete this line

	return &cc.KvListReply{}, nil
}
