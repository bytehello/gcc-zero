package releasekv

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFilePutContents(t *testing.T) {
	configPath, _ := filepath.Abs("config")
	filename := "baojia/test.conf"
	svr := NewService(configPath)
	err := svr.FilePutContents(filename, "test_value")
	if err != nil {
		t.Fatalf("file put contents err: %s", err.Error())
	}
	os.RemoveAll(configPath)
}
