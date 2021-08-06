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

type ClusterListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClusterListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ClusterListLogic {
	return ClusterListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClusterListLogic) ClusterList(req types.ClusterListReq) (*types.ClusterListReply, error) {
	var listReq ccclient.ClusterListReq
	_ = copier.Copy(&listReq, &req)
	reply, err := l.svcCtx.CcRpcClient.ClusterList(l.ctx, &listReq)
	if err != nil {
		return nil, utilsErr.ConvertErrorx(err)
	}
	var list []*types.ClusterListData
	for _, v := range reply.Data {
		var temp types.ClusterListData
		_ = copier.Copy(&temp, &v)
		list = append(list, &temp)
	}
	return &types.ClusterListReply{
		Code:     "0",
		Message:  "success",
		Data:     list,
		Current:  req.Current,
		PageSize: req.PageSize,
		Total:    reply.Total,
	}, nil
}
