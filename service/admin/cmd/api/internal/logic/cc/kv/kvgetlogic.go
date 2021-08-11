package logic

import (
	"context"

	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/svc"
	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type KvGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKvGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) KvGetLogic {
	return KvGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KvGetLogic) KvGet(req types.KvGetReq) (*types.KvGetReply, error) {
	// todo: add your logic here and delete this line

	return &types.KvGetReply{}, nil
}
