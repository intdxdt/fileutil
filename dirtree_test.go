package fileutil

import (
	"os"
	"testing"
	"path/filepath"
	"github.com/franela/goblin"
	"fmt"
)

func TestTree(t *testing.T) {
	pwd, _ := os.Getwd()
	path := filepath.Join(pwd, "test")
	g := goblin.Goblin(t)
	g.Describe("stats", func() {
		g.It("dir stats", func() {
			s, err := NewStat("test/01")
			g.Assert(err == nil).IsTrue()
			g.Assert(s != nil).IsTrue()
			g.Assert(s.IsFile()).IsFalse()
			g.Assert(s.IsDir()).IsTrue()
			g.Assert(s.Path() == "test/01").IsTrue()
			fmt.Println(s.ModTime())
			fmt.Println(s.ModTimeAsString())
			fmt.Println(s.Size())
			s, err = NewStat("test/x01")
			g.Assert(err != nil).IsTrue()
			g.Assert(s == nil).IsTrue()

		})
	})
	g.Describe("directory tree", func() {
		g.It("dir tree", func() {
			ignoreDirs := []string{".bzr", "hg", ".git", ".idea", ".path"}
			res, err := Tree(path, ignoreDirs)
			g.Assert(len(res)).Equal(11)
			g.Assert(err == nil).IsTrue()
		})
	})
}
