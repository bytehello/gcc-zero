package logic

import (
	"context"

	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/cc"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type KvReleaseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewKvReleaseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KvReleaseLogic {
	return &KvReleaseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *KvReleaseLogic) KvRelease(in *cc.KvReleaseReq) (*cc.KvReleaseReply, error) {
	// todo: add your logic here and delete this line

	return &cc.KvReleaseReply{}, nil
}
