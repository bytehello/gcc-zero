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

type ClusterDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClusterDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) ClusterDelLogic {
	return ClusterDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClusterDelLogic) ClusterDel(req types.ClusterDelReq) (*types.ClusterDelReply, error) {
	var delReq ccclient.ClusterDelReq
	_ = copier.Copy(&delReq, &req)
	_, err := l.svcCtx.CcRpcClient.ClusterDel(l.ctx, &delReq)
	if err != nil {
		return nil, utilsErr.ConvertErrorx(err)
	}
	return &types.ClusterDelReply{
		Code:    "0",
		Message: "success",
	}, nil
}
