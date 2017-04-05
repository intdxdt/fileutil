package fileutil

import (
	"os"
	"testing"
	"path/filepath"
	"github.com/franela/goblin"
)

func TestTree(t *testing.T) {
	pwd, _ := os.Getwd()
	path :=filepath.Join(pwd, "test")
	g := goblin.Goblin(t)
	g.Describe("directory tree", func() {
		g.It("dir tree", func() {
			ignoreDirs := []string{".bzr", "hg", ".git", ".idea", ".path"}
			res, err := Tree(path, ignoreDirs)
			 g.Assert(len(res)).Equal(11)
			 g.Assert(err == nil ).IsTrue()
		})
	})
}
