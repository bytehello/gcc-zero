package logic

import (
	"context"
	"github.com/bytehello/gcc-zero/common/errorx"

	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/cc"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type AppDelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAppDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppDelLogic {
	return &AppDelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AppDelLogic) AppDel(in *cc.AppDelReq) (*cc.AppDelReply, error) {
	err := l.svcCtx.AppModel.Delete(in.Id)
	if err != nil {
		return nil, errorx.DefaultCodeError(err.Error())
	}
	return &cc.AppDelReply{}, nil
}
