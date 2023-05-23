// A group of functions about file system and native api
package sys

import (
	"fmt"
	"os"
	"os/signal"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
	"unicode"
)

const (
	LINUX   = "linux"
	WINDOWS = "windows"
	DARWIN  = "darwin"
	UNKNOWN = "unknown"
)

func OS() string {
	return runtime.GOOS
}

func IsWindows() bool {
	return OS() == WINDOWS
}
func IsLinux() bool {
	return OS() == LINUX
}

func Platform() string {
	return runtime.GOOS + "-" + runtime.GOARCH
}

func IsDarwin() bool {
	return OS() == DARWIN
}

// get GOPATH from env
func GoPath() (dir string, ok bool) {
	dir, ok = os.LookupEnv("GOPATH")
	return
}

// More complete file stat struct
type FileStat struct {
	Name   string
	Abs    string
	Base   string
	Parent string
	Seps   []string
	Info   os.FileInfo
}

// Parse path and return new FileStat
func NewFStat(path string) (fstat *FileStat, err error) {
	stat, err := os.Stat(path)
	if err != nil {
		return
	}
	abspath, _ := filepath.Abs(path)
	fstat = &FileStat{
		path,
		abspath,
		filepath.Base(path),
		filepath.Dir(abspath),
		strings.SplitAfter(abspath, string(os.PathSeparator)),
		stat,
	}
	return
}

// FileSize get file size in bytes
func FileSize(path string) int64 {
	stat, err := os.Stat(path)
	if err != nil {
		return -1
	}
	return stat.Size()
}

// Walk in dir and return all file's FileStat
func ListDir(path string) []FileStat {
	var list = make([]FileStat, 0, 10)
	path, _ = filepath.Abs(path)
	filepath.Walk(path, func(path string, _ os.FileInfo, err error) error {
		if err == nil {
			if filepath, _ := NewFStat(path); filepath != nil {
				list = append(list, *filepath)
			}
		}
		return err
	})
	return list
}

// Walk in dir and call user function func(path string, isdir bool)
func ScanDir(name string, callback func(path string, isdir bool)) {
	filepath.Walk(name, func(path string, info os.FileInfo, err error) error {
		callback(path, info.IsDir())
		return err
	})
}

// Get current user's home dir
func HomeDir() (string, bool) {
	if IsWindows() {
		u, err := user.Current()
		if err == nil {
			return u.HomeDir, true
		}
		return "", false
	}
	return os.LookupEnv("HOME")
}

// ExpandHome expand ~ to user's home dir
func ExpandHome(path string) string {
	if len(path) == 0 || path[0] != '~' {
		return path
	}
	u, _ := user.Current()
	return u.HomeDir + path[1:]
}

// ExpandAbs expand path to absolute path
func ExpandAbs(path string) string {
	path, _ = filepath.Abs(ExpandHome(path))
	return path
}

// ProgramDir return dir of program use os.Args[0]
func ProgramDir() string {
	path, _ := filepath.Abs(os.Args[0])
	return filepath.Dir(path)
}

// Return parent dir name of path
func DirName(path string) string {
	return filepath.Dir(path)
}

// Return parent dir name of path
func ParentName(path string) string {
	path, _ = filepath.Abs(path)
	return filepath.Dir(path)
}

// Return grand parent dir name of path
func GrandName(path string) string {
	path, _ = filepath.Abs(path)
	return filepath.Dir(filepath.Dir(path))
}

// IsWinRoot check whether a path is windows absolute path with disk letter
func IsWinRoot(path string) bool {
	if path == "" {
		return false
	}
	return unicode.IsLetter(rune(path[0])) && strings.HasPrefix(path[1:], ":\\")
}

// IsRoot check wether or not path is root of filesystem
func IsUnixRoot(path string) bool {
	switch OS() {
	case WINDOWS:
		return IsWinRoot(path)
	case LINUX, DARWIN:
		return path == "/"
	default:
		return false
	}
}

func RemoveExt(path string) string {
	pos := strings.LastIndex(path, ".")
	if pos > 0 {
		return path[:pos]
	}
	return path
}

func ReplaceExt(path, ext string) string {
	if strings.HasPrefix(ext, ".") {
		return RemoveExt(path) + ext
	}
	return RemoveExt(path) + "." + ext
}

// FileExist check whether or not file exist
func FileExist(path string) bool {
	stat, err := os.Stat(path)
	return err == nil && !stat.IsDir()
}

// DirExist check whether or not given name is a dir
func DirExist(path string) bool {
	stat, err := os.Stat(path)
	return err == nil && stat.IsDir()
}

// IsSymlink check whether or not given name is a symlink
func IsSymlink(fname string) bool {
	_, err := os.Lstat(fname)
	return err == nil
}

// a simple signal notifier
func RegistSignalHandler(handler func(os.Signal), signals ...os.Signal) {
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, signals...)
		s := <-c
		handler(s)
	}()
}

// Return stack list for printing
func StackInfo(full bool) string {
	var buf [4096]byte
	n := runtime.Stack(buf[:], full)
	return string(buf[:n])
}

// Print stack info
func PrintStack() {
	fmt.Fprintf(os.Stderr, "%s\n", StackInfo(false))
}

// Call this only in main once
func RunTimeInit(max_num_threads int) {
	if max_num_threads == 0 {
		if runtime.NumCPU() == 1 {
			runtime.GOMAXPROCS(1)
		} else {
			runtime.GOMAXPROCS(runtime.NumCPU()/2 + 1)
		}
	} else {
		runtime.GOMAXPROCS(max_num_threads)
	}
}
