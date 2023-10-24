// A group of functions about file system and native api
package system

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
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

func ChangeWorkDir(dir string, verbose bool) error {
	cwd, _ := os.Getwd()
	if cwd == dir {
		return nil
	}
	err := os.Chdir(dir)
	if err != nil {
		return err
	}
	cwd, _ = os.Getwd()
	if verbose {
		fmt.Fprintf(os.Stderr, "Working directory: %s\n", cwd)
	}
	return nil
}
