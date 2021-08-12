package logic

import (
	"context"
	utilsErr "github.com/bytehello/gcc-zero/common/grpc/utils/err"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/ccclient"
	"github.com/jinzhu/copier"

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
	var rpcReq ccclient.KvUpdateReq
	_ = copier.Copy(&rpcReq, req)
	_, err := l.svcCtx.CcRpcClient.KvUpdate(l.ctx, &rpcReq)
	if err != nil {
		return nil, utilsErr.ConvertErrorx(err)
	}
	return &types.KvUpdateReply{
		Code: "0",
	}, nil
}
