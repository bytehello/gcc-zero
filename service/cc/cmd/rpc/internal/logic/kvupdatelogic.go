package logic

import (
	"context"
	"errors"
	"github.com/bytehello/gcc-zero/internal/bizerror"
	"github.com/bytehello/gcc-zero/service/cc/model/ccmodel"

	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/cc"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type KvUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewKvUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KvUpdateLogic {
	return &KvUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *KvUpdateLogic) KvUpdate(in *cc.KvUpdateReq) (*cc.KvUpdateReply, error) {
	kv, err := l.svcCtx.KvModel.FindOne(in.Id)
	if err != nil {
		return nil, bizerror.New(bizerror.ErrCodeKvFind)
	}
	tempKv, err := l.svcCtx.KvModel.FindOneByAppIdClusterIdKey(kv.AppId, kv.ClusterId, in.Key)
	if err != nil {
		if !errors.Is(err, ccmodel.ErrNotFound) {
			return nil, bizerror.New(bizerror.ErrCodeKvFind)
		}
	} else {
		if tempKv.Id != in.Id {
			return nil, bizerror.New(bizerror.ErrCodeKvAddKeyExisted)
		}
	}
	err = l.svcCtx.KvModel.UpdateKeyValueDesc(ccmodel.CcKv{
		Id:    in.Id,
		Key:   in.Key,
		Value: in.Value,
		Desc:  in.Desc,
	})
	if err != nil {
		return nil, bizerror.Newf(bizerror.ErrCodeKvAddKeyUpdate, "KvUpdate err:%s", err.Error())
	}
	return &cc.KvUpdateReply{}, nil
}
