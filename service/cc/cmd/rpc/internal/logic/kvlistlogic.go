package logic

import (
	"context"

	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/cc"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type KvListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewKvListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KvListLogic {
	return &KvListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *KvListLogic) KvList(in *cc.KvListReq) (*cc.KvListReply, error) {
	all := l.svcCtx.KvModel.FindAll(in.AppId, in.ClusterId, in.Current, in.PageSize)
	total := l.svcCtx.KvModel.Count(in.AppId, in.ClusterId)
	var list []*cc.KvData
	for _, v := range *all {
		list = append(list, &cc.KvData{
			Id:             v.Id,
			AppId:          v.AppId,
			ClusterId:      v.ClusterId,
			Key:            v.Key,
			Value:          v.Value,
			Desc:           v.Desc,
			Version:        v.Version,
			PushStatus:     v.PushStatus,
			Format:         v.Format,
			CreateRevision: v.CreateRevision,
			ModRevision:    v.ModRevision,
			CreateTime:     v.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime:     v.UpdateTime.Format("2006-01-02 15:04:05"),
			// TODO sql null
			//DeletedTime:    v.DeletedTime.Format("2006-01-02 15:04:05"),
			//PushedTime:     v.PushedTime.Format("2006-01-02 15:04:05"),
		})
	}
	return &cc.KvListReply{
		Total:    total,
		PageSize: in.PageSize,
		Current:  in.Current,
		List:     list,
	}, nil
}
