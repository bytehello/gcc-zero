package logic

import (
	"context"

	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/svc"
	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type KvClientListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKvClientListLogic(ctx context.Context, svcCtx *svc.ServiceContext) KvClientListLogic {
	return KvClientListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KvClientListLogic) KvClientList(req types.KvClientListReq) (*types.KvClientListReply, error) {
	// todo: add your logic here and delete this line

	return &types.KvClientListReply{}, nil
}
