package file

import (
	"path/filepath"
	"testing"
)

func TestRealPath(t *testing.T) {
	path := "file.go"
	abs, err := filepath.Abs(path)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(abs)
	t.Log(Exists(path))
	t.Log(Dir(path))
}
