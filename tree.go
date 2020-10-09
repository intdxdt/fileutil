package fileutil

import (
	"os"
	"path/filepath"
)

func filter(results *[]string, ignoreDirs []string) filepath.WalkFunc {
	ignoreMap := make(map[string]bool)
	for _, d := range ignoreDirs {
		ignoreMap[d] = true
	}

	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			key := filepath.Base(path)
			if ignoreMap[key] {
				return filepath.SkipDir
			}
			*results = append(*results, path)
		}
		return nil
	}
}

func Tree(directory string, ignoreDirs []string) ([]string, error) {
	results := make([]string, 0)
	err := filepath.Walk(directory, filter(&results, ignoreDirs))
	return results, err
}
