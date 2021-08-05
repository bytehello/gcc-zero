package logic

import (
	"context"
	"github.com/bytehello/gcc-zero/common/errorx"
	"github.com/bytehello/gcc-zero/internal/bizerror"
	"github.com/bytehello/gcc-zero/service/cc/cmd/model/ccmodel"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/cc"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type AppAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAppAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppAddLogic {
	return &AppAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AppAddLogic) AppAdd(in *cc.AppAddReq) (*cc.AppAddReply, error) {
	res, err := l.svcCtx.AppModel.Insert(ccmodel.CcApp{
		Desc:    in.Desc,
		AppKey:  in.AppKey,
		AppName: in.AppName,
	})
	if err != nil {
		return nil, bizerror.Customized(1, "插入失败")
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, errorx.DefaultCodeError(err.Error())
	}
	return &cc.AppAddReply{Id: id}, nil
}
