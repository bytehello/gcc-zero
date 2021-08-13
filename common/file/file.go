package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	// Separator for file system.
	// It here defines the separator as variable
	// to allow it modified by developer if necessary.
	Separator = string(filepath.Separator)

	// DefaultPerm is the default perm for file opening.
	DefaultPermOpen = os.FileMode(0666)

	// DefaultPermCopy is the default perm for file/folder copy.
	DefaultPermCopy = os.FileMode(0777)

	// selfPath is the current running binary path.
	// As it is most commonly used, it is so defined as an internal package variable.
	selfPath = ""

	// Temporary directory of system.
	tempDir = "/tmp"
)

func Exists(path string) bool {
	if s, err := os.Stat(path); s != nil && !os.IsNotExist(err) {
		return true
	}
	return false
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// Dir returns all but the last element of path, typically the path's directory.
// After dropping the final element, Dir calls Clean on the path and trailing
// slashes are removed.
// If the `path` is empty, Dir returns ".".
// If the `path` is ".", Dir treats the path as current working directory.
// If the `path` consists entirely of separators, Dir returns a single separator.
// The returned path does not end in a separator unless it is the root directory.
func Dir(path string) string {
	if path == "." {
		return filepath.Dir(RealPath(path))
	}
	return filepath.Dir(path)
}

// RealPath converts the given <path> to its absolute path
// and checks if the file path exists.
// If the file does not exist, return an empty string.
func RealPath(path string) string {
	p, err := filepath.Abs(path)
	if err != nil {
		return ""
	}
	if !Exists(p) {
		return ""
	}
	return p
}

// Create creates file with given <path> recursively.
// The parameter <path> is suggested to be absolute path.
func Create(path string) (*os.File, error) {
	dir := Dir(path)
	if !Exists(dir) {
		if err := Mkdir(dir); err != nil {
			return nil, err
		}
	}
	return os.Create(path)
}

// Mkdir creates directories recursively with given <path>.
// The parameter <path> is suggested to be an absolute path instead of relative one.
func Mkdir(path string) error {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}
	return nil
}

// Remove deletes all file/directory with <path> parameter.
// If parameter <path> is directory, it deletes it recursively.
func Remove(path string) error {
	return os.RemoveAll(path)
}

// IsWritable checks whether given <path> is writable.
//
// TODO improve performance; use golang.org/x/sys to cross-plat-form
func IsWritable(path string) bool {
	result := true
	if IsDir(path) {
		// If it's a directory, create a temporary file to test whether it's writable.
		tmpFile := strings.TrimRight(path, Separator) + Separator + fmt.Sprintf("%d", time.Now().UnixNano())
		if f, err := Create(tmpFile); err != nil || !Exists(tmpFile) {
			result = false
		} else {
			f.Close()
			Remove(tmpFile)
		}
	} else {
		// 如果是文件，那么判断文件是否可打开
		file, err := os.OpenFile(path, os.O_WRONLY, DefaultPermOpen)
		if err != nil {
			result = false
		}
		file.Close()
	}
	return result
}
