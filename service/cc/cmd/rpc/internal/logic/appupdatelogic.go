package logic

import (
	"context"
	"github.com/bytehello/gcc-zero/common/errorx"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/cc"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/internal/svc"
	"github.com/bytehello/gcc-zero/service/cc/model/ccmodel"

	"github.com/tal-tech/go-zero/core/logx"
)

type AppUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAppUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppUpdateLogic {
	return &AppUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AppUpdateLogic) AppUpdate(in *cc.AppUpdateReq) (*cc.AppUpdateReply, error) {
	err := l.svcCtx.AppModel.Update(ccmodel.CcApp{
		Id:      in.Id,
		Desc:    in.Desc,
		AppKey:  in.AppKey,
		AppName: in.AppName,
	})
	if err != nil {
		return nil, errorx.DefaultCodeError(err.Error())
	}
	return &cc.AppUpdateReply{}, nil
}
