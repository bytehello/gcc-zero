package logic

import (
	"context"
	"github.com/bytehello/gcc-zero/internal/bizerror"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/cc"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/ccclient"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type KvClientListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewKvClientListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KvClientListLogic {
	return &KvClientListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *KvClientListLogic) KvClientList(in *cc.KvClientListReq) (*cc.KvClientListReply, error) {
	list, err := l.svcCtx.ClientModel.FindAllByKvId(in.KvId)
	if err != nil {
		return nil, bizerror.Newf(bizerror.ErrCodeClientFind, "KvClientList err:", err.Error())
	}
	var result []*ccclient.KvClientData
	for _, v := range list {

		result = append(result, &ccclient.KvClientData{
			Ip:          v.Ip,
			VisitedTime: v.VisitedTime.Time.Format("2006-01-02 15:04:05"),
			CreateTime:  v.CreateTime.Format("2006-01-02 15:04:05"),
			UpdateTime:  v.UpdateTime.Format("2006-01-02 15:04:05"),
			ReleaseTime: v.ReleaseTime.Time.Format("2006-01-02 15:04:05"),
			KvId:        v.KvId,
			AppId:       v.AppId,
			ClusterId:   v.ClusterId,
		})
	}
	return &cc.KvClientListReply{List: result}, nil
}
