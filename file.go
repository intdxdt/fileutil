package fileutil

import (
	"os"
	"time"
	"fmt"
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

//new stat
func NewStat(fname string) (*Stat, error) {
	info, err := os.Stat(fname)
	if os.IsNotExist(err) {
		return nil, err
	}

	t := isFile
	if info.Mode().IsDir() {
		t = isDir
	}

	return &Stat{
		path:    fname,
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

//is file
func IsFile(fname string) bool {
	info, err := os.Stat(fname)
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

//modified time
func ModifiedTime(fname string) (time.Time, error) {
	info, err := os.Stat(fname)
	if os.IsNotExist(err) {
		return time.Time{}, err
	}
	return info.ModTime(), nil
}

//modified time as unix string
func ModUnixString(fname string) (string, error) {
	var updated int64
	mt, err := ModifiedTime(fname)
	if err == nil {
		updated = mt.Unix()
	}
	return fmt.Sprintf("%v", updated), err
}

//Are all paths directories
func AreAllDir(dirs []string) bool {
	bln := true
	for i := 0; bln && i < len(dirs); i++ {
		bln = bln && IsDir(dirs[i])
	}
	return bln
}
