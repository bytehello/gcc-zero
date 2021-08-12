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
	getReq := ccclient.KvGetReq{}
	_ = copier.Copy(&getReq, &req)
	resp, err := l.svcCtx.CcRpcClient.KvGet(l.ctx, &getReq)
	if err != nil {
		return nil, utilsErr.ConvertErrorx(err)
	}
	kvData := types.KvData{}
	_ = copier.Copy(&kvData, resp.Data)
	return &types.KvGetReply{
		Code:    "0",
		Message: "success",
		Data:    kvData,
	}, nil
}
