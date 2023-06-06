package fio

import (
	"os"
	"strings"
	"testing"

	"github.com/alexloser/goaux/fs"
)

func TestFio(t *testing.T) {
	name := fs.GrandName(fs.ExpandAbs(os.Args[len(os.Args)-1])) + string(os.PathSeparator) + "README.md"
	if !fs.FileExist(name) {
		t.Logf(name)
		t.Error(name)
		return
	}

	first, err := FirstLine(name)
	t.Log(first)
	if err != nil {
		t.Error(err)
		return
	}

	lines, _ := ReadLines(name)
	if lines[0] != first {
		t.Error(lines[0])
		t.Error(first)
	}

	f, r, _ := MakeReader(name)
	defer f.Close()

	if line, _ := r.ReadString('\n'); strings.Trim(line, "\n") != first {
		t.Error(line)
	}

	var last1, last2 string
	var n1, n2 int

	ReadBytesLine(name, func(line []byte) {
		n1++
		last1 = string(line)
	})

	ReadLine(name, func(line string) {
		n2++
		last2 = line
	})

	if n1 != n2 || last2 != last1 {
		t.Fail()
	}
}
