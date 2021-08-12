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
	var rpcReq ccclient.KvDelReq
	_ = copier.Copy(&rpcReq, req)
	_, err := l.svcCtx.CcRpcClient.KvDel(l.ctx, &rpcReq)
	if err != nil {
		return nil, utilsErr.ConvertErrorx(err)
	}
	return &types.KvDelReply{
		Code:    "0",
		Message: "success",
	}, nil
}
