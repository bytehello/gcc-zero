package logic

import (
	"context"
	"github.com/bytehello/gcc-zero/common/errorx"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/cc"
	"github.com/jinzhu/copier"

	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/svc"
	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AppDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAppDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) AppDelLogic {
	return AppDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AppDelLogic) AppDel(req types.AppDelReq) (*types.AppDelReply, error) {
	var appDelReq cc.AppDelReq
	_ = copier.Copy(&appDelReq, req)
	_, err := l.svcCtx.CcRpcClient.AppDel(l.ctx, &appDelReq)
	if err != nil {
		return nil, errorx.DefaultCodeError(err.Error())
	}
	return &types.AppDelReply{
		Code:    "0",
		Message: "success",
	}, nil
}
