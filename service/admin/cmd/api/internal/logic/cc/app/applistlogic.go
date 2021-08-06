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

type AppListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAppListLogic(ctx context.Context, svcCtx *svc.ServiceContext) AppListLogic {
	return AppListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AppListLogic) AppList(req types.AppListReq) (*types.AppListReply, error) {
	appListReq := ccclient.AppListReq{}
	_ = copier.Copy(&appListReq, &req)
	applistReply, err := l.svcCtx.CcRpcClient.AppList(l.ctx, &appListReq)
	if err != nil {
		// 1。直接返回错误交给http error 处理
		// 2。返回正常的Reply
		// 这里选择1
		return nil, utilsErr.ConvertErrorx(err)
	}
	// TODO rpc 返回的list 和 api 的 list 不一致，如何处理
	var appList []*types.ListAppData
	for _, v := range applistReply.List {
		temp := types.ListAppData{}
		_ = copier.Copy(&temp, &v)
		appList = append(appList, &temp)
	}
	return &types.AppListReply{
		Code:     "0",
		Message:  "success",
		Data:     appList,
		Current:  req.Current,
		PageSize: req.PageSize,
		Total:    applistReply.Total,
	}, nil
}
