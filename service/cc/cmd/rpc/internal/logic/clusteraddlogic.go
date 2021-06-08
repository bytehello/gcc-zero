package logic

import (
	"context"
	"github.com/bytehello/gcc-zero/common/errorx"
	"github.com/bytehello/gcc-zero/service/cc/cmd/model/ccmodel"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/cc"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type ClusterAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClusterAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClusterAddLogic {
	return &ClusterAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClusterAddLogic) ClusterAdd(in *cc.ClusterAddReq) (*cc.ClusterAddReply, error) {
	_, err := l.svcCtx.AppModel.FindOne(in.AppId)
	if err != nil {
		return nil, errorx.DefaultCodeError(err.Error())
	}
	res, err := l.svcCtx.ClusterModel.Insert(ccmodel.CcCluster{
		ClusterName: in.ClusterName,
		Desc:        in.Desc,
		AppId:       in.AppId,
	})
	if err != nil {
		return nil, errorx.DefaultCodeError(err.Error())
	}
	var id int64
	id, err = res.LastInsertId()
	if err != nil {
		return nil, errorx.DefaultCodeError(err.Error())
	}

	return &cc.ClusterAddReply{
		Id: id,
	}, nil
}
