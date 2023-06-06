package system

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/alexloser/goaux/fs"
)

func TestSystem(t *testing.T) {
	path, ok := os.LookupEnv("GOROOT")
	if !ok {
		t.Fatal("GOROOT")
	}
	if fs.DirExist(path) != true {
		t.Fatal(path)
	}

	fstat, err := fs.NewFStat(path)
	if err != nil {
		t.Fatal(err)
	}
	abs, _ := filepath.Abs(path)
	if fstat.Abs != abs {
		t.Fatal(path)
	}
	if fstat.Base != filepath.Base(path) {
		t.Fatal(path)
	}
	if fstat.Parent != filepath.Dir(path) {
		t.Fatal(path)
	}
	if len(fstat.Seps) < 1 {
		t.Fatal(path)
	}

	ret := fs.ListDir(path)
	if len(ret) == 0 {
		t.Error(path)
	}

	fs.ScanDir(path, func(name string, isdir bool) {
		if (isdir && !fs.DirExist(name)) || (!isdir && fs.DirExist(name)) {
			t.Error(name)
		}
	})

	if fs.RemoveExt("bar.txt") != "bar" {
		t.Fail()
	}
	if fs.ReplaceExt("foo.csv", "dat") != "foo.dat" {
		t.Fail()
	}

	if IsLinux() {
		if fs.DirName("/home/user/check/data") != "/home/user/check" {
			t.Fail()
		}
	}
	if IsWindows() {
		if fs.DirName("C:/data/temp") != "C:\\data" {
			t.Fail()
		}
	}
}
