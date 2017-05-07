package fileutil

import (
	"os"
	"time"
)

//is file
func IsFile(fname string) bool {
	info, err := os.Stat(fname)
	if os.IsNotExist(err) {
		return false
	}
	return info.Mode().IsRegular()
}

//modified time
func ModifiedTime(fname string) (time.Time, error) {
	info, err := os.Stat(fname)
	if os.IsNotExist(err) {
		return time.Time{}, err
	}
	return info.ModTime(), nil
}

func IsDir(dirname string) bool {
	info, err := os.Stat(dirname)
	bln := false
	if err == nil && info.Mode().IsDir() {
		bln = true
	}
	return bln
}

//Are all paths directories
func AreAllDir(dirs []string) bool {
	bln := true
	for i := 0; bln && i < len(dirs); i++ {
		bln = bln && IsDir(dirs[i])
	}
	return bln
}
