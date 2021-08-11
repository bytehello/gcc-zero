package logic

import (
	"context"
	"errors"
	"github.com/bytehello/gcc-zero/internal/bizerror"
	"github.com/bytehello/gcc-zero/service/cc/cmd/model/ccmodel"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/cc"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type ClusterUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClusterUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClusterUpdateLogic {
	return &ClusterUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClusterUpdateLogic) ClusterUpdate(in *cc.ClusterUpdateReq) (*cc.ClusterUpdateReply, error) {
	one, err := l.svcCtx.ClusterModel.FindOne(in.Id)
	if err != nil {
		return nil, bizerror.New(bizerror.ErrCodeClusterFind)
	}
	cluster, err := l.svcCtx.ClusterModel.GetCluster(one.AppId, in.ClusterName)
	if err != nil {
		if !errors.Is(err, ccmodel.ErrNotFound) {
			return nil, bizerror.New(bizerror.ErrCodeClusterFind)
		}
	} else if cluster.Id != in.Id { // 有重复
		return nil, bizerror.Newf(bizerror.ErrCodeCLusterNameExisted, "existed cluster id: %d, app id: %s", cluster.Id, cluster.AppId)
	}
	err = l.svcCtx.ClusterModel.Update(ccmodel.CcCluster{
		Id:          in.Id,
		ClusterName: in.ClusterName,
		Desc:        in.Desc,
	})
	if err != nil {
		return nil, bizerror.New(bizerror.ErrCodeClusterUpdate)
	}
	return &cc.ClusterUpdateReply{}, nil
}
