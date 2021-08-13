package logic

import (
	"context"
	"database/sql"
	"github.com/bytehello/gcc-zero/internal"
	"github.com/bytehello/gcc-zero/service/cc/model/ccmodel"
	"github.com/pkg/errors"
	"time"

	"github.com/bytehello/gcc-zero/internal/bizerror"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/cc"
	"github.com/bytehello/gcc-zero/service/cc/cmd/rpc/internal/svc"
	"github.com/tal-tech/go-zero/core/logx"
)

const (
	ErrCodeKvAddClusterNotFound   = iota + 1001001 // cluster 未查询到
	ErrCodeKvAddClusterFind                        // cluster 查询失败
	ErrCodeClusterIdNotMatchAppId                  // 未匹配
	ErrCodeKvAddParamsInvalid                      // 参数有误
	ErrCodeKvAddKeyFind                            // 查找 key 失败
	ErrCodeKvAddKeyExisted                         // key 已经存在
	ErrCodeKvAddKeyInsert                          // 插入 key 失败
	ErrCodeKvAddEtcdPut                            // 存入 etcd 失败
	ErrCodeKvAddEtcdGet                            // 读取etcd 失败
	ErrCodeKvAddKeyUpdate                          // 更新 kv 失败
)

type KvAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewKvAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KvAddLogic {
	return &KvAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *KvAddLogic) KvAdd(in *cc.KvAddReq) (*cc.KvAddReply, error) {
	var (
		cluster *ccmodel.CcCluster
		err     error
	)
	if in.Value == "" || in.Key == "" || in.AppId == 0 || in.ClusterId == 0 {
		l.Logger.Error("KvAdd 参数不合法：", in)
		return nil, bizerror.New(bizerror.ErrCodeDefaultBadParams)
	}

	if cluster, err = l.svcCtx.ClusterModel.FindOne(in.ClusterId); err != nil {
		if errors.Is(err, ccmodel.ErrNotFound) {
			return nil, bizerror.New(bizerror.ErrCodeKvAddClusterNotFound)
		}
		return nil, bizerror.New(bizerror.ErrCodeKvAddClusterFind)
	}
	if cluster.AppId != in.AppId {
		return nil, bizerror.New(bizerror.ErrCodeClusterIdNotMatchAppId)
	}

	if _, err = l.svcCtx.KvModel.FindOneByAppIdClusterIdKey(in.AppId, in.ClusterId, in.Key); err != nil {
		if !errors.Is(err, ccmodel.ErrNotFound) {
			return nil, bizerror.New(bizerror.ErrCodeKvAddKeyFind)
		}
	} else {
		return nil, bizerror.New(bizerror.ErrCodeKvAddKeyExisted)
	}
	// TODO 开启事务

	kvData := ccmodel.CcKv{
		Key:        in.Key,
		ClusterId:  in.ClusterId,
		AppId:      in.AppId,
		Value:      in.Value,
		Format:     internal.ValueFormatJson,
		CreateTime: time.Now(),
		Desc:       in.Desc,
	}
	var result sql.Result
	if result, err = l.svcCtx.KvModel.Insert(kvData); err != nil {
		return nil, bizerror.Newf(bizerror.ErrCodeKvAddKeyInsert, "保存失败, %s", err.Error())
	}
	// etcd client 操作
	var keyValue *internal.KeyValue
	if _, err = l.svcCtx.KVer.Put(in.Key, in.Value); err != nil {
		l.Logger.Error("KvAdd etcd put err:", in.Key, in.Value, err)
		return nil, bizerror.Newf(bizerror.ErrCodeKvAddEtcdPut, "etcd 保存失败, %s", err.Error())
	}
	id, _ := result.LastInsertId()
	if keyValue, err = l.svcCtx.KVer.Get(in.Key); err != nil {
		l.Logger.Error("KvAdd etcd get err:", in.Key, err)
		return nil, bizerror.Newf(bizerror.ErrCodeKvAddEtcdGet, "etcd 读取失败, %s", err.Error())
	}
	kvData.Id = id
	kvData.Version = keyValue.Version
	kvData.CreateRevision = keyValue.CreateRevision
	kvData.ModRevision = keyValue.ModRevision
	if err = l.svcCtx.KvModel.Update(kvData); err != nil {
		return nil, bizerror.Newf(bizerror.ErrCodeKvAddKeyUpdate, "更新 kv 失败, %s", err.Error())
	}
	return &cc.KvAddReply{
		Id:             kvData.Id,
		Version:        kvData.Version,
		CreateRevision: kvData.CreateRevision,
		ModRevision:    kvData.ModRevision,
	}, nil
}
