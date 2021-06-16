package logic

import (
	"context"

	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/cc"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type KvClientListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewKvClientListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KvClientListLogic {
	return &KvClientListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *KvClientListLogic) KvClientList(in *cc.KvClientListReq) (*cc.KvClientListReply, error) {
	// todo: add your logic here and delete this line

	return &cc.KvClientListReply{}, nil
}
