package fio

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

// Capacity in bytes
const (
	_ = 1 << (10 * iota)
	KB
	MB
	GB
	TB
	PB
)

// Define default linebreak char and io buffer size
const (
	IO_BUF_SIZE = 8192
	LINE_END    = byte('\n')
	BOM         = "\xEF\xBB\xBF"
)

// MakeReader Use defer file.Close() after using this if err is nil
func MakeReader(path string) (*os.File, *bufio.Reader, error) {
	file, err := os.Open(path)
	if err != nil {
		return file, nil, err
	}
	reader := bufio.NewReaderSize(file, 8192)
	return file, reader, err
}

// ReadBytesLine Read file by line, need a callback accepted bytes of line
func ReadBytesLine(path string, callback func(bstr []byte)) (uint, error) {
	file, reader, err := MakeReader(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	var nlines uint = 0
	for {
		bs, err := reader.ReadSlice(LINE_END)
		switch err {
		case nil:
			nlines++
			if len(bs) > 0 {
				callback(bs[0 : len(bs)-1]) // strip
			} else {
				callback(bs)
			}
		case io.EOF:
			return nlines, nil
		default:
			return nlines, err
		}
	}
}

// ReadLine Read file by line, need a callback accepted string pointer of line
func ReadLine(path string, callback func(str string)) (uint, error) {
	file, reader, err := MakeReader(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	var nlines uint = 0
	for {
		str, err := reader.ReadString(LINE_END)
		switch err {
		case nil:
			nlines++
			if len(str) > 0 {
				callback(str[0 : len(str)-1]) // strip
			} else {
				callback(str)
			}
		case io.EOF:
			return nlines, nil
		default:
			return nlines, err
		}
	}

}

// FirstLine Get first line of text file
func FirstLine(path string) (line string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		line = scanner.Text()
	}

	return
}

// ReadLines Get all lines in text file
func ReadLines(path string) ([]string, error) {
	content, err := os.ReadFile(path)

	if err != nil {
		return nil, err
	}

	return strings.Split(string(content), string(LINE_END)), nil
}

func HasBOM(path string) bool {
	if handle, err := os.Open(path); err == nil {
		defer handle.Close()
		buf := make([]byte, 3)
		n, _ := handle.Read(buf)
		if n == 3 {
			return BOM == string(buf)
		}
	}
	return false
}

func WriteFile(path string, data interface{}) error {
	file, _ := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	defer func() {
		file.Sync()
		file.Close()
	}()
	switch data.(type) {
	case string:
		file.WriteString(data.(string))
		return nil
	case []byte:
		file.Write(data.([]byte))
		return nil
	default:
		return errors.New("Invalid data type")
	}
}
