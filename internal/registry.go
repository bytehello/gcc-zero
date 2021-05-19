package internal

import (
	"github.com/tal-tech/go-zero/core/syncx"
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

func (r *Registry) GetConn() (EtcdClient, error) {
	return nil, nil
}

func (r *Registry) GetCluster(endpoints []string) *cluster {
	r.lock.Lock()
	defer r.lock.Unlock()
	key := getClusterKey(endpoints)
	if cluster, ok := r.clusters[key]; ok {
		return cluster
	} else {
		//core/discov/internal/registry.go
	}
	return nil
}

func getClusterKey(endpoints []string) string {
	sort.Strings(endpoints)
	return strings.Join(endpoints, endpointsSeparator)
}

type cluster struct {
	endpoints []string
	key       string
}
