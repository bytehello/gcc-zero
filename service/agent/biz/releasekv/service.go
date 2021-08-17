package releasekv

import (
	"fmt"
	"github.com/bytehello/gcc-zero/common/file"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/core/syncx"
	"github.com/tal-tech/go-zero/zrpc"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

const (
	rpcClientPrefix = "gcc-zero:agent:rpc:client:"
)

type Service struct {
	releasePath string
	resource    map[string]zrpc.Client
	shareCall   syncx.SharedCalls
	lock        sync.RWMutex
}

func NewService(path string) *Service {
	if !file.Exists(path) {
		if err := file.Mkdir(path); err != nil {
			panic(fmt.Sprintf("new release service err: mkdir config path err : %s", err.Error()))
		}
	}
	return &Service{
		releasePath: path,
		resource:    make(map[string]zrpc.Client),
		shareCall:   syncx.NewSharedCalls(),
	}
}

func (s *Service) FilePutContents(filename string, value string) error {
	realFilename := s.realConfigFileName(filename)
	if !file.Exists(realFilename) {
		if _, err := file.Create(realFilename); err != nil {
			return err
		}
	}
	if err := ioutil.WriteFile(realFilename, []byte(value), file.DefaultPermOpen); err != nil {
		return err
	}
	return nil
}

func (s *Service) ConfigFileStat(filename string) (os.FileInfo, error) {
	return os.Stat(s.realConfigFileName(filename))
}

func (s *Service) realConfigFileName(filename string) string {
	return filepath.Join(s.releasePath, filename)
}

// getClient 获取 zrpc 客户端
func (s *Service) getClient(target string) (zrpc.Client, error) {
	return s.getResource(s.getClientKey(target), func() (zrpc.Client, error) {
		logx.Info("Init zrpc client:", target)
		return zrpc.NewClientWithTarget(target)
	})
}

func (s *Service) getClientKey(target string) string {
	return fmt.Sprintf("%s%s", rpcClientPrefix, target)
}

func (s *Service) getResource(key string, create func() (zrpc.Client, error)) (zrpc.Client, error) {
	val, err := s.shareCall.Do(key, func() (interface{}, error) {
		s.lock.RLock()
		client, ok := s.resource[key]
		s.lock.RUnlock()
		if ok {
			return client, nil
		}
		val, err := create()
		if err != nil {
			return nil, err
		}

		s.lock.Lock()
		s.resource[key] = val
		s.lock.Unlock()
		return val, err
	})
	if err != nil {
		return nil, err
	}
	return val.(zrpc.Client), nil
}
