// Text test
package text

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

func TestText(t *testing.T) {
	a := []int{1, 2, 3}
	s := ToString(a)
	if s != fmt.Sprintf("%v", a) {
		t.Fail()
	}

	s = ToSyntaxString(a)
	if s != fmt.Sprintf("%#v", a) {
		t.Fail()
	}

	s = "A#B#C#D#E"

	r1, ok1 := SSplit(s, "#", 3)
	t.Log(ToSyntaxString(r1))

	r2, ok2 := BSplit([]byte(s), []byte("#"), 3)
	t.Log(ToSyntaxString(r2))

	ret := StringSliceFilter(r1, func(s string) bool { return strings.Contains(s, "#") })
	t.Log(ToSyntaxString(ret))

	if len(r1) != 3 && ok1 {
		t.Fail()
	}
	if len(r2) != 3 && ok2 {
		t.Fail()
	}
	if IndexOf("B", r1) != 1 {
		t.Fail()
	}

	s = TrimBlanks(" \vgood boy\t\r\n")
	if s != "good boy" {
		t.Fatal(s)
	}

	f32 := float32(3232.3232)
	f64 := float64(640640640.646464)

	i32 := int32(320320)
	i64 := int64(640640640)

	if Atoi32("320320") != i32 || Itoa32(i32) != "320320" {
		t.Fail()
	}
	if Atoi64("640640640") != i64 || Itoa64(i64) != "640640640" {
		t.Fail()
	}
	if math.Abs(float64(Atof32("3232.3232")-f32)) > 1e-6 || Ftoa32(f32) != "3232.3232" {
		t.Fail()
	}

	if math.Abs(Atof64("640640640.646464")-f64) > 1e-6 || Ftoa64(f64) != "6.40640640646464E+08" {
		t.Log(math.Abs(Atof64("640640640.646464") - f64))
		t.Fail()
	}

	t.Log(ToString(f32))
	t.Log(ToString(f64))
}
