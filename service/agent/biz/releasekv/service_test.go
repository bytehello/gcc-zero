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

func TestGetClient(t *testing.T) {
	configPath, _ := filepath.Abs("config")
	s := NewService(configPath)
	for i := 0; i < 10; i++ {
		_, err := s.getClient("127.0.0.1:8080")
		if err != nil {
			t.Fatalf("Get client err: %s", err.Error())
		}
	}
}
