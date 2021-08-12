package logic

import (
	"context"
	utilsErr "github.com/bytehello/gcc-zero/common/grpc/utils/err"
	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/svc"
	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/types"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/ccclient"
	"github.com/tal-tech/go-zero/core/logx"
)

type KvClientListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKvClientListLogic(ctx context.Context, svcCtx *svc.ServiceContext) KvClientListLogic {
	return KvClientListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KvClientListLogic) KvClientList(req types.KvClientListReq) (*types.KvClientListReply, error) {
	var rpcReq ccclient.KvClientListReq
	rpcReq.KvId = req.KvId
	reply, err := l.svcCtx.CcRpcClient.KvClientList(l.ctx, &rpcReq)
	if err != nil {
		return nil, utilsErr.ConvertErrorx(err)
	}
	var result []*types.KvClientData
	if len(reply.List) > 0 {
		for _, v := range reply.List {
			result = append(result, &types.KvClientData{
				Ip:          v.Ip,
				VisitedTime: v.VisitedTime,
				CreateTime:  v.CreateTime,
				UpdateTime:  v.UpdateTime,
				ReleaseTime: v.ReleaseTime,
				KvId:        v.KvId,
				AppId:       v.AppId,
				ClusterId:   v.ClusterId,
			})
		}
	}
	return &types.KvClientListReply{List: result}, nil
}
