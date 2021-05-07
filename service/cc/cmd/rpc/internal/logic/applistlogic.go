package logic

import (
	"context"

	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/cc"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type AppListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAppListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppListLogic {
	return &AppListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AppListLogic) AppList(in *cc.AppListReq) (*cc.AppListReply, error) {
	all, _ := l.svcCtx.AppModel.FindAll(in.Current, in.PageSize)
	total, _ := l.svcCtx.AppModel.Count()
	var list []*cc.ListAppData
	for _, v := range *all {
		list = append(list, &cc.ListAppData{
			Id:         v.Id,
			AppKey:     v.AppKey,
			AppName:    v.AppName,
			Desc:       v.Desc,
			CreateTime: v.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime: v.UpdateTime.Format("2006-01-02 15:04:05"),
		})
	}
	return &cc.AppListReply{
		Total: total,
		List:  list,
	}, nil
}
