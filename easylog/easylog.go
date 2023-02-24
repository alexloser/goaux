// An easy using logger for Go, thread-safe with high performance
package easylog

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Default global logger
var elog *EasyLog = nil

// Only 5 log levels
const (
	D int = iota // Debug
	I int = iota // Info
	W int = iota // Warning
	E int = iota // Error
	F int = iota // Fatal or Final
)

// Internal lock for IO
var once_flag sync.Once

// For better performance in singel thread
type dummy_mutex struct{}

func (m *dummy_mutex) Lock() {
	// do nothing
}

func (m *dummy_mutex) Unlock() {
	// do nothing
}

// This simple logger has only one member io.Writer for writing data
type EasyLog struct {
	mtx   sync.Locker
	file  *os.File
	buf   bytes.Buffer
	level int
}

// Init global logger, if unsafe is true, dummy mutex will be used instead of sync.Mutex
func Initialize(file *os.File, unsafe ...bool) *EasyLog {
	if elog != nil {
		Warn("EasyLog has already initialized!!!")
	}

	once_flag.Do(func() {
		if len(unsafe) > 0 && unsafe[0] {
			elog = &EasyLog{&dummy_mutex{}, file, bytes.Buffer{}, I}
		} else {
			elog = &EasyLog{&sync.Mutex{}, file, bytes.Buffer{}, I}
		}
	})

	return elog
}

// Set log level
func SetLevel(level int) {
	elog.mtx.Lock()
	defer elog.mtx.Unlock()
	elog.level = level
}

// Get log level
func GetLevel() int {
	return elog.level
}

// Sync all data in buffer into opened log-file
func Flush() {
	elog.mtx.Lock()
	defer elog.mtx.Unlock()
	elog.file.Sync()
}

// Join all log prefix
func Join(buf *bytes.Buffer, level int, name string, line int) string {
	buf.Reset()
	buf.WriteString(time.Now().String()[:23])

	switch level {
	case D:
		buf.WriteString(" D [")
	case I:
		buf.WriteString(" I [")
	case W:
		buf.WriteString(" W [")
	case E:
		buf.WriteString(" E [")
	case F:
		buf.WriteString(" F [")
	}

	buf.WriteString(strconv.Itoa(os.Getpid()))
	buf.WriteByte(' ')
	buf.WriteString(name)
	buf.WriteByte(':')
	buf.WriteString(strconv.Itoa(line))
	buf.WriteString("] ")

	return buf.String()
}

// Standard log style printer, it's multithread safe, output format like:
// 2019-04-12 18:01:29.244 I [6460 main.go:62] this is a info
func Log(file *os.File, degree int, level int, vargs ...interface{}) {
	elog.mtx.Lock()
	defer elog.mtx.Unlock()
	if level < elog.level {
		return
	}

	_, name, line, _ := runtime.Caller(degree)

	if pos := strings.LastIndex(name, "/"); pos != -1 {
		name = name[pos+1:]
		if strings.HasSuffix(name, ".go") {
			name = strings.Replace(name, ".go", "", 1)
		}
	}

	fmt.Fprintf(file, Join(&elog.buf, level, name, line))

	for _, v := range vargs {
		fmt.Fprintf(file, "%v ", v)
	}

	fmt.Fprintf(file, "\n")
}

// On level info
func Info(v ...interface{}) {
	Log(elog.file, 2, I, v...)
}

// On level debug
func Debug(v ...interface{}) {
	Log(elog.file, 2, D, v...)
}

// On level warning
func Warn(v ...interface{}) {
	Log(elog.file, 2, W, v...)
}

// On level error
func Error(v ...interface{}) {
	Log(elog.file, 2, E, v...)
}

// On level fatal
func Fatal(v ...interface{}) {
	Log(elog.file, 2, F, v...)
}

// Safe assert(means exit when false)
func Assert(condition bool) {
	if false == condition {
		Log(os.Stderr, 2, E, "Assert Failed!")
		os.Exit(-1)
	}
}

// UnSafe assert(means no exit when false)
func AssertNoExit(condition bool) {
	if false == condition {
		Log(os.Stderr, 2, E, "Assert Failed!")
	}
}
