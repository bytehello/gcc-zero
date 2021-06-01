package internal

import (
	"github.com/pkg/errors"
	"github.com/tal-tech/go-zero/core/syncx"
)

var ErrKeyNotFound = errors.New("etcd key not found")

type KVer struct {
	endpoints []string
	quit      *syncx.DoneChan
}

type KeyValue struct {
	Key            []byte
	Value          []byte
	CreateRevision int64
	ModRevision    int64
	Version        int64
}

func NewKVer(endpoints []string) *KVer {
	return &KVer{
		endpoints: endpoints,
		quit:      syncx.NewDoneChan(),
	}
}

func (kv *KVer) Get(key string) (*KeyValue, error) {
	cli, err := GetRegistry().GetConn(kv.endpoints)
	if err != nil {
		return nil, nil
	}
	resp, err := cli.Get(cli.Ctx(), key)
	if err != nil {
		return nil, err
	}
	if len(resp.Kvs) == 0 {
		return nil, errors.Wrap(ErrKeyNotFound, "not found")
	}
	return &KeyValue{
		Key:            resp.Kvs[0].Key,
		Value:          resp.Kvs[0].Value,
		CreateRevision: resp.Kvs[0].CreateRevision,
		ModRevision:    resp.Kvs[0].ModRevision,
		Version:        resp.Kvs[0].Version,
	}, nil
}

func (kv *KVer) Put(key string, value string) (prevKeyValue *KeyValue, err error) {
	var cli EtcdClient
	cli, err = GetRegistry().GetConn(kv.endpoints)
	if err != nil {
		return nil, errors.Wrap(err, "etcd getConn fail")
	}
	putResp, err := cli.Put(cli.Ctx(), key, value)
	if err != nil {
		return nil, errors.Wrap(err, "etcd put fail")
	}
	if putResp.PrevKv == nil {
		return
	}
	prevKeyValue = &KeyValue{
		Key:            putResp.PrevKv.Key,
		Value:          putResp.PrevKv.Value,
		CreateRevision: putResp.PrevKv.CreateRevision,
		ModRevision:    putResp.PrevKv.ModRevision,
		Version:        putResp.PrevKv.Version,
	}
	return
}

func (kv *KVer) Stop() {
	kv.quit.Close()
}
