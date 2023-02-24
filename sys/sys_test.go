package sys

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSystem(t *testing.T) {
	path, ok := os.LookupEnv("GOROOT")
	if !ok {
		t.Fatal("GOROOT")
	}
	if DirExist(path) != true {
		t.Fatal(path)
	}

	fs, err := NewFStat(path)
	if err != nil {
		t.Fatal(err)
	}
	abs, _ := filepath.Abs(path)
	if fs.Abs != abs {
		t.Fatal(path)
	}
	if fs.Base != filepath.Base(path) {
		t.Fatal(path)
	}
	if fs.Parent != filepath.Dir(path) {
		t.Fatal(path)
	}
	if len(fs.Seps) < 1 {
		t.Fatal(path)
	}

	ret := ListDir(path)
	if len(ret) == 0 {
		t.Error(path)
	}

	ScanDir(path, func(name string, isdir bool) {
		if (isdir && !DirExist(name)) || (!isdir && DirExist(name)) {
			t.Error(name)
		}
	})

	if RemoveExt("bar.txt") != "bar" {
		t.Fail()
	}
	if ReplaceExt("foo.csv", "dat") != "foo.dat" {
		t.Fail()
	}

	if IsLinux() {
		if DirName("/home/user/check/data") != "/home/user/check" {
			t.Fail()
		}
	}
	if IsWindows() {
		if DirName("C:/data/temp") != "C:\\data" {
			t.Fail()
		}
	}
}
