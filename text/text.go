// Text
package text

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
)

// String to int64 ignore error
func Atoi32(s string) int32 {
	n, _ := strconv.ParseInt(s, 10, 32)
	return int32(n)
}

// String to int64 ignore error
func Atoi64(s string) int64 {
	n, _ := strconv.ParseInt(s, 10, 64)
	return n
}

// String to int64 ignore error
func Atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

// Int64 to string
func Itoa32(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}

// Int64 to string
func Itoa64(i int64) string {
	return strconv.FormatInt(i, 10)
}

// Int to string
func Itoa(i int) string {
	return strconv.Itoa(i)
}

// String to int64 ignore error
func Atof32(s string) float32 {
	f, _ := strconv.ParseFloat(s, 32)
	return float32(f)
}

// String to int64 ignore error
func Atof64(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

// Int64 to string
func Ftoa32(f float32) string {
	return strconv.FormatFloat(float64(f), 'G', -1, 32)
}

// Int64 to string
func Ftoa64(f float64) string {
	return strconv.FormatFloat(f, 'G', -1, 64)
}

// To string for everything
func ToString(val interface{}) string {
	return fmt.Sprintf("%v", val)
}

// To syntax string for everything
func ToSyntaxString(val interface{}) string {
	return fmt.Sprintf("%#v", val)
}

// Split string and check length of results
func SSplit(str string, delim string, expect int) (seq []string, ok bool) {
	seq = strings.SplitN(str, delim, expect)
	ok = len(seq) == expect
	return
}

// Split bytes and check length of results
func BSplit(str []byte, delim []byte, expect int) (seq [][]byte, ok bool) {
	seq = bytes.SplitN(str, delim, expect)
	ok = len(seq) == expect
	return
}

// Get index of value in the array
func IndexOf(value string, seq []string) int {
	for idx, item := range seq {
		if item == value {
			return idx
		}
	}
	return -1
}

// Trim all blank chars of string
func TrimBlanks(s string) string {
	return strings.Trim(s, " \t\v\r\n\\0")
}

// Trim all blank chars of string only on left
func TrimLeftBlanks(s string) string {
	return strings.TrimLeft(s, " \t\v\r\n\\0")
}

// Trim all blank chars of string only on right
func TrimRightBlanks(s string) string {
	return strings.TrimRight(s, " \t\v\r\n\\0")
}

// Filter for each string in strings slice
func StringSliceFilter(slice []string, filter func(string) bool) []string {
	ret := make([]string, 0, len(slice)/2)
	for _, s := range slice {
		if filter(s) {
			ret = append(ret, s)
		}
	}
	return ret
}

func GBK2UTF8(src string) (string, error) {
	utf, err := simplifiedchinese.GBK.NewDecoder().String(src)
	if err != nil {
		return "", err
	}
	return utf, err
}
