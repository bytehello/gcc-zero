package internal

import (
	"github.com/tal-tech/go-zero/core/syncx"
	"go.etcd.io/etcd/clientv3"
	"io"
	"sort"
	"strings"
	"sync"
)

// 1.声明遍历，指针
// 2。lock用法
var (
	registry = Registry{
		clusters: make(map[string]*cluster),
	}
	connManager = syncx.NewResourceManager()
)

type Registry struct {
	clusters map[string]*cluster
	lock     sync.Mutex
}

func GetRegistry() *Registry {
	return &registry
}

func (r *Registry) GetConn(endpoints []string) (EtcdClient, error) {
	return r.GetCluster(endpoints).getClient()
}

func (r *Registry) GetCluster(endpoints []string) *cluster {
	key := getClusterKey(endpoints)
	r.lock.Lock()
	defer r.lock.Unlock()
	if cluster, ok := r.clusters[key]; ok {
		return cluster
	} else {
		//refer: core/discov/internal/registry.go
		c := newCluster(endpoints)
		r.clusters[key] = c
		return c
	}
}

func getClusterKey(endpoints []string) string {
	sort.Strings(endpoints)
	return strings.Join(endpoints, endpointsSeparator)
}

func newCluster(endpoints []string) *cluster {
	return &cluster{
		endpoints: endpoints,
		key:       getClusterKey(endpoints),
	}
}

type cluster struct {
	endpoints []string
	key       string
}

func (c *cluster) getClient() (EtcdClient, error) {
	val, err := connManager.GetResource(c.key, func() (closer io.Closer, err error) {
		return c.newClient()
	})
	return val.(EtcdClient), err
}

func (c *cluster) newClient() (EtcdClient, error) {
	v, err := clientv3.New(clientv3.Config{
		Endpoints:            c.endpoints,
		AutoSyncInterval:     autoSyncInterval,
		DialTimeout:          dialTimeout,
		DialKeepAliveTime:    dialKeepAliveTime,
		DialKeepAliveTimeout: dialTimeout,
		RejectOldCluster:     true,
	})
	if err != nil {
		return nil, err
	}
	return v, nil
}
