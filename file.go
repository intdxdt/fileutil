package fileutil

import (
	"fmt"
	"os"
	"time"
)

const (
	isFile = -1
	isDir  = 1
)

type Stat struct {
	path    string
	size    int64
	modTime int64
	_type   int
}

func NewStat(filename string) (*Stat, error) {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return &Stat{}, err
	}

	t := isFile
	if info.Mode().IsDir() {
		t = isDir
	}

	return &Stat{
		path:    filename,
		size:    info.Size(),
		modTime: info.ModTime().Unix(),
		_type:   t,
	}, nil
}

func (s *Stat) IsFile() bool {
	return s._type == isFile
}

func (s *Stat) IsDir() bool {
	return s._type == isDir
}

func (s *Stat) Path() string {
	return s.path
}
func (s *Stat) Size() int64 {
	return s.size
}
func (s *Stat) ModTime() int64 {
	return s.modTime
}

func (s *Stat) ModTimeAsString() string {
	return fmt.Sprintf("%v", s.ModTime())
}

func IsFile(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return info.Mode().IsRegular()
}

func IsDir(dirname string) bool {
	info, err := os.Stat(dirname)
	bln := false
	if err == nil && info.Mode().IsDir() {
		bln = true
	}
	return bln
}

func ModifiedTime(filename string) (time.Time, error) {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return time.Time{}, err
	}
	return info.ModTime(), nil
}

func ModUnixString(filename string) (string, error) {
	var updated int64
	mt, err := ModifiedTime(filename)
	if err == nil {
		updated = mt.Unix()
	}
	return fmt.Sprintf("%v", updated), err
}

func AreAllDir(dirs []string) bool {
	bln := true
	for i := 0; bln && i < len(dirs); i++ {
		bln = bln && IsDir(dirs[i])
	}
	return bln
}
