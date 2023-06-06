// An easy using logger for Go, thread-safe with high performance
package easylog

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestEasyLog(t *testing.T) {
	var log = Initialize(os.Stdout, true)
	SetLevel(0)

	Debug("debug", "foo")
	Debug("debug", "BAR")
	Info("info", "foo")
	Info("info", "bar")
	Warn("warn", "foo")
	Warn("warn", "bar")
	Error("error", "foo")
	Error("error", "bar")

	fmt.Println()
	SetLevel(2)
	Debug("debug", "no")
	Info("info", "no")
	Warn("warn", "foo")
	Error("error", "bar")

	Assert(0 == D)
	Assert(1 == I)
	Assert(2 == W)
	Assert(3 == E)
	Assert(4 == F)

	var dup = Initialize(os.Stdout, true)
	AssertNoExit(log == dup)

	b := bytes.Buffer{}
	fmt.Printf("%s OK\n", Join(&b, 1, "easylog_test.go", 32))

}
