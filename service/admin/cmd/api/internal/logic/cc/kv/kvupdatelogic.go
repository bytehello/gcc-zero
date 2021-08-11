package logic

import (
	"context"

	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/svc"
	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type KvUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKvUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) KvUpdateLogic {
	return KvUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KvUpdateLogic) KvUpdate(req types.KvUpdateReq) (*types.KvUpdateReply, error) {
	// todo: add your logic here and delete this line

	return &types.KvUpdateReply{}, nil
}
