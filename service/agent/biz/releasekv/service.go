package releasekv

import (
	"fmt"
	"github.com/bytehello/gcc-zero/common/file"
	"os"
	"path/filepath"
)

type Service struct {
	releasePath string
}

func NewService(path string) *Service {
	if !file.Exists(path) {
		if err := file.Mkdir(path); err != nil {
			panic(fmt.Sprintf("new release service err: mkdir config path err : %s", err.Error()))
		}
	}
	return &Service{
		releasePath: path,
	}
}

func (svr *Service) FilePutContents(filename string, value string) error {
	realFilename := svr.realConfigFileName(filename)
	if !file.Exists(realFilename) {
		if _, err := file.Create(realFilename); err != nil {
			return err
		}
	}
	if err := os.WriteFile(realFilename, []byte(value), file.DefaultPermOpen); err != nil {
		return err
	}
	return nil
}

func (svr *Service) ConfigFileStat(filename string) (os.FileInfo, error) {
	return os.Stat(svr.realConfigFileName(filename))
}

func (svr *Service) realConfigFileName(filename string) string {
	return filepath.Join(svr.releasePath, filename)
}
