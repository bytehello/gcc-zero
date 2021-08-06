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

type AppAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAppAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) AppAddLogic {
	return AppAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AppAddLogic) AppAdd(req types.AppAddReq) (*types.AppAddReply, error) {
	var addReq ccclient.AppAddReq
	_ = copier.Copy(&addReq, req)
	reply, err := l.svcCtx.CcRpcClient.AppAdd(l.ctx, &addReq)
	if err != nil {
		return nil, utilsErr.ConvertErrorx(err)
	}
	return &types.AppAddReply{
		Id:      reply.Id,
		Code:    "0",
		Message: "success",
	}, nil
}
