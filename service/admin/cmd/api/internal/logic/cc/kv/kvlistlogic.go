package logic

import (
	"context"
	utilsErr "github.com/bytehello/gcc-zero/common/grpc/utils/err"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/ccclient"

	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/svc"
	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type KvListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKvListLogic(ctx context.Context, svcCtx *svc.ServiceContext) KvListLogic {
	return KvListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KvListLogic) KvList(req types.KvListReq) (*types.KvListReply, error) {
	rpcResp, err := l.svcCtx.CcRpcClient.KvList(l.ctx, &ccclient.KvListReq{
		AppId:     req.AppId,
		ClusterId: req.ClusterId,
		Current:   req.Current,
		PageSize:  req.PageSize,
	})
	if err != nil {
		return nil, utilsErr.ConvertErrorx(err)
	}
	var list []*types.KvData
	if len(rpcResp.List) > 0 {
		for _, v := range rpcResp.List {
			list = append(list, &types.KvData{
				Id:             v.Id,
				AppId:          v.AppId,
				ClusterId:      v.ClusterId,
				Key:            v.Key,
				Value:          v.Value,
				Desc:           v.Desc,
				Version:        v.Version,
				PushStatus:     int(v.PushStatus),
				Format:         v.Format,
				CreateRevision: v.CreateRevision,
				ModRevision:    v.ModRevision,
				CreateTime:     v.CreateTime,
				UpdateTime:     v.UpdateTime,
				PushedTime:     v.PushedTime,
			})
		}
	}
	return &types.KvListReply{
		Code:     "0",
		Message:  "success",
		Total:    rpcResp.Total,
		PageSize: rpcResp.PageSize,
		Current:  rpcResp.Current,
		List:     list,
	}, nil
}
