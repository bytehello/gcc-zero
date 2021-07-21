package logic

import (
	"context"
	"github.com/bytehello/gcc-zero/common/api/protobuf"
	"github.com/bytehello/gcc-zero/common/errorx"
	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/svc"
	"github.com/bytehello/gcc-zero/service/admin/cmd/api/internal/types"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/ccclient"
	"github.com/jinzhu/copier"
	"github.com/tal-tech/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type KvAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKvAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) KvAddLogic {
	return KvAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KvAddLogic) KvAdd(req types.KvAddReq) (*types.KvAddReply, error) {
	var (
		addReq   ccclient.KvAddReq
		addReply *ccclient.KvAddReply
		err      error
	)
	_ = copier.Copy(&addReq, &req)
	l.Logger.Infof("KvAdd req:", addReq)
	if addReply, err = l.svcCtx.CcRpcClient.KvAdd(l.ctx, &addReq); err != nil {
		s := status.Convert(err)
		d := s.Details()[0]
		switch info := d.(type) {
		case *protobuf.BizError:
			return nil, errorx.NewCodeError(info.ErrCode, info.ErrMsg)
		default:
			return nil, errorx.DefaultCodeError("未知错误")
		}
	}
	return &types.KvAddReply{
		Code:           "0",
		Id:             addReply.Id,
		Version:        addReply.Version,
		CreateRevision: addReply.CreateRevision,
		ModRevision:    addReply.ModRevision,
		Message:        "success",
	}, nil
}
