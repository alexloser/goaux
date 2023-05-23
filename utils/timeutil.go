// Date and time utils
package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var Weekdays = []string{
	"Sunday",
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
}

var ShortWeekdays = []string{
	"Sun",
	"Mon",
	"Tue",
	"Wed",
	"Thu",
	"Fri",
	"Sat",
}

var ShortMonths = []string{
	"---",
	"Jan",
	"Feb",
	"Mar",
	"Apr",
	"May",
	"Jun",
	"Jul",
	"Aug",
	"Sep",
	"Oct",
	"Nov",
	"Dec",
}

var Months = []string{
	"---",
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

func weekNumber(t *time.Time, char int) int {
	weekday := int(t.Weekday())

	if char == 'W' {
		// Monday as the first day of the week
		if weekday == 0 {
			weekday = 6
		} else {
			weekday -= 1
		}
	}

	return (t.YearDay() + 6 - weekday) / 7
}

func SecondToTime(sec int64) time.Time {
	return time.Unix(sec, 0)
}

// Strftime formats time.Date according to the directives in the given format string.
// The directives begins with a percent (%) character.
func Strftime(t *time.Time, f string) string {
	var result []string
	format := []rune(f)

	add := func(str string) {
		result = append(result, str)
	}

	for i := 0; i < len(format); i++ {
		switch format[i] {
		case '%':
			if i < len(format)-1 {
				switch format[i+1] {
				case 'a':
					add(ShortWeekdays[t.Weekday()])
				case 'A':
					add(Weekdays[t.Weekday()])
				case 'w':
					add(fmt.Sprintf("%d", t.Weekday()))
				case 'd':
					add(fmt.Sprintf("%02d", t.Day()))
				case 'b':
					add(ShortMonths[t.Month()])
				case 'B':
					add(Months[t.Month()])
				case 'm':
					add(fmt.Sprintf("%02d", t.Month()))
				case 'y':
					add(fmt.Sprintf("%02d", t.Year()%100))
				case 'Y':
					add(fmt.Sprintf("%02d", t.Year()))
				case 'H':
					add(fmt.Sprintf("%02d", t.Hour()))
				case 'I':
					if t.Hour() == 0 {
						add(fmt.Sprintf("%02d", 12))
					} else if t.Hour() > 12 {
						add(fmt.Sprintf("%02d", t.Hour()-12))
					} else {
						add(fmt.Sprintf("%02d", t.Hour()))
					}
				case 'p':
					if t.Hour() < 12 {
						add("AM")
					} else {
						add("PM")
					}
				case 'M':
					add(fmt.Sprintf("%02d", t.Minute()))
				case 'S':
					add(fmt.Sprintf("%02d", t.Second()))
				case 'f':
					add(fmt.Sprintf("%06d", t.Nanosecond()/1000))
				case 'z':
					add(t.Format("-0700"))
				case 'Z':
					add(t.Format("MST"))
				case 'j':
					add(fmt.Sprintf("%03d", t.YearDay()))
				case 'U':
					add(fmt.Sprintf("%02d", weekNumber(t, 'U')))
				case 'W':
					add(fmt.Sprintf("%02d", weekNumber(t, 'W')))
				case 'c':
					add(t.Format("Mon Jan 2 15:04:05 2006"))
				case 'x':
					add(fmt.Sprintf("%02d/%02d/%02d", t.Month(), t.Day(), t.Year()%100))
				case 'X':
					add(fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second()))
				case '%':
					add("%")
				}
				i += 1
			}
		default:
			add(string(format[i]))
		}
	}

	return strings.Join(result, "")
}

func abs(v time.Duration) time.Duration {
	if v < 0 {
		v *= -1
	}
	return v
}

// Timedelta represents a duration between two dates.
// All fields are optional and default to 0. You can initialize any type of timedelta by specifying field values which you want to use.
type Timedelta struct {
	Days, Seconds, Microseconds, Milliseconds, Minutes, Hours, Weeks time.Duration
}

// Add returns the Timedelta t+t2.
func (t *Timedelta) Add(t2 *Timedelta) Timedelta {
	return Timedelta{
		Days:         t.Days + t2.Days,
		Seconds:      t.Seconds + t2.Seconds,
		Microseconds: t.Microseconds + t2.Microseconds,
		Milliseconds: t.Milliseconds + t2.Milliseconds,
		Minutes:      t.Minutes + t2.Minutes,
		Hours:        t.Hours + t2.Hours,
		Weeks:        t.Weeks + t2.Weeks,
	}
}

// Subtract returns the Timedelta t-t2.
func (t *Timedelta) Subtract(t2 *Timedelta) Timedelta {
	return Timedelta{
		Days:         t.Days - t2.Days,
		Seconds:      t.Seconds - t2.Seconds,
		Microseconds: t.Microseconds - t2.Microseconds,
		Milliseconds: t.Milliseconds - t2.Milliseconds,
		Minutes:      t.Minutes - t2.Minutes,
		Hours:        t.Hours - t2.Hours,
		Weeks:        t.Weeks - t2.Weeks,
	}
}

// Abs returns the absolute value of t
func (t *Timedelta) Abs() Timedelta {
	return Timedelta{
		Days:         abs(t.Days),
		Seconds:      abs(t.Seconds),
		Microseconds: abs(t.Microseconds),
		Milliseconds: abs(t.Milliseconds),
		Minutes:      abs(t.Minutes),
		Hours:        abs(t.Hours),
		Weeks:        abs(t.Weeks),
	}
}

// Duration() returns time.Duration. time.Duration can be added to time.Date.
func (t *Timedelta) Duration() time.Duration {
	return t.Days*24*time.Hour +
		t.Seconds*time.Second +
		t.Microseconds*time.Microsecond +
		t.Milliseconds*time.Millisecond +
		t.Minutes*time.Minute +
		t.Hours*time.Hour +
		t.Weeks*7*24*time.Hour
}

// String returns a string representing the Timedelta's duration in the form "72h3m0.5s".
func (t *Timedelta) String() string {
	return t.Duration().String()
}

// Wait(to sleep) until hour:minute:second arrived
func WaitUntil(hour, minute, second int) {
	if hour < -1 || hour >= 24 {
		panic("hour is not correct")
	}

	if minute < -1 || minute >= 60 {
		panic("minute is not correct")
	}

	if second < -1 || second >= 60 {
		panic("second is not correct")
	}

	var check = func() bool {
		now := time.Now().String()[11:19]
		t := strings.Split(now[:8], ":")

		h, _ := strconv.Atoi(t[0])
		m, _ := strconv.Atoi(t[1])
		s, _ := strconv.Atoi(t[2])

		if hour != -1 && h != hour {
			return false
		}
		if minute != -1 && m != minute {
			return false
		}
		if second != -1 && s != second {
			return false
		}

		return true
	}

	for {
		if check() {
			return
		}
		time.Sleep(time.Millisecond * 500)
	}
}
