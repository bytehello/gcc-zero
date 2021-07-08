package logic

import (
	"context"
	"github.com/bytehello/gcc-zero/common/errorx"

	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/cc"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type ClusterListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClusterListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClusterListLogic {
	return &ClusterListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClusterListLogic) ClusterList(in *cc.ClusterListReq) (*cc.ClusterListReply, error) {
	list, err := l.svcCtx.ClusterModel.FindAll(in.Current, in.PageSize)
	total, _ := l.svcCtx.ClusterModel.Count()
	if err != nil {
		return nil, errorx.DefaultCodeError(err.Error())
	}
	var data []*cc.ClusterListData
	for _, v := range *list {
		data = append(data, &cc.ClusterListData{
			ClusterName: v.ClusterName,
			Desc:        v.Desc,
			Id:          v.Id,
			AppId:       v.AppId,
		})
	}
	return &cc.ClusterListReply{
		Data:  data,
		Total: total,
	}, nil
}
