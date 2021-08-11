package logic

import (
	"context"

	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/svc"
	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type KvDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKvDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) KvDelLogic {
	return KvDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KvDelLogic) KvDel(req types.KvDelReq) (*types.KvDelReply, error) {
	// todo: add your logic here and delete this line

	return &types.KvDelReply{}, nil
}
