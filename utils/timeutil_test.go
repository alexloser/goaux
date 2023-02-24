package utils

import (
	"reflect"
	"testing"
	"time"
)

func TestStrftime(t *testing.T) {
	date := time.Date(2005, 2, 3, 4, 5, 6, 7000, time.UTC)
	AssertEqual(t, Strftime(&date,
		"%a %A %w %d %b %B %m %y %Y %H %I %p %M %S %f %z %Z %j %U %W %c %x %X %%"),
		"Thu Thursday 4 03 Feb February 02 05 2005 04 04 AM 05 06 000007 +0000 UTC 034 05 05 Thu Feb 3 04:05:06 2005 02/03/05 04:05:06 %")

	date = time.Date(2015, 7, 2, 15, 24, 30, 35, time.UTC)
	AssertEqual(t, Strftime(&date, "%U %W"), "26 26")

	date = time.Date(1962, 3, 23, 15, 24, 30, 35, time.UTC)
	AssertEqual(t, Strftime(&date, "%U %W"), "11 12")

	date = time.Date(1989, 12, 31, 15, 24, 30, 35000, time.UTC)
	AssertEqual(t, Strftime(&date, "%U %W"), "53 52")

	AssertEqual(t, Strftime(&date,
		"%a %A %w %d %b %B %m %y %Y %H %I %p %M %S %f %z %Z %j %U %W %c %x %X %%"),
		"Sun Sunday 0 31 Dec December 12 89 1989 15 03 PM 24 30 000035 +0000 UTC 365 53 52 Sun Dec 31 15:24:30 1989 12/31/89 15:24:30 %")

	date = time.Date(1989, 12, 31, 0, 24, 30, 35000, time.UTC)
	AssertEqual(t, Strftime(&date, "%I"), "12")

	AssertEqual(t, Strftime(&date, "%a %A %w %d %b %B %"), "Sun Sunday 0 31 Dec December ")

	AssertEqual(t, Strftime(&date, "%a %A %w %d %b %B %"), "Sun Sunday 0 31 Dec December ")
}

func AssertEqual(t *testing.T, x, y interface{}) {
	if !reflect.DeepEqual(x, y) {
		t.Error("Expected ", y, ", got ", x)
	}
}

func TestTimedelta(t *testing.T) {
	base := time.Date(1980, 1, 6, 0, 0, 0, 0, time.UTC)
	result := base.Add((&Timedelta{Days: 1, Seconds: 66355, Weeks: 1722}).Duration())
	AssertEqual(t, result.String(), "2013-01-07 18:25:55 +0000 UTC")

	result = result.Add((&Timedelta{Microseconds: 3, Milliseconds: 10, Minutes: 1}).Duration())
	AssertEqual(t, result.String(), "2013-01-07 18:26:55.010003 +0000 UTC")

	td := Timedelta{Days: 10, Minutes: 17, Seconds: 56}
	td2 := Timedelta{Days: 15, Minutes: 13, Seconds: 42}
	td = td.Add(&td2)

	base = time.Date(2015, 2, 3, 0, 0, 0, 0, time.UTC)
	result = base.Add(td.Duration())
	AssertEqual(t, result.String(), "2015-02-28 00:31:38 +0000 UTC")

	td = td.Subtract(&td2)

	result = base.Add(td.Duration())
	AssertEqual(t, result.String(), "2015-02-13 00:17:56 +0000 UTC")

	AssertEqual(t, td.String(), "240h17m56s")

	td = Timedelta{Days: -1, Seconds: -1, Microseconds: -1, Milliseconds: -1, Minutes: -1, Hours: -1, Weeks: -1}
	td2 = td
	td = td.Abs()
	td = td.Add(&td2)
	AssertEqual(t, td.String(), "0s")
}
