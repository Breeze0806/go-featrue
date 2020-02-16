package others

import (
	"bytes"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestBytes_ToValidUTF8(t *testing.T) {
	for _, v := range toValidUTF8Tests {
		out := bytes.ToValidUTF8([]byte(v.input.s), []byte(v.input.replacement))
		if !Equal(out, []byte(v.want)) {
			t.Errorf("ToValidUTF8(%q, %q) = %q; want %q", v.input.s, v.input.replacement, out, v.want)
		}
	}
}

func TestStrings_ToValidUTF8(t *testing.T) {
	for _, v := range toValidUTF8Tests {
		out := strings.ToValidUTF8(v.input.s, v.input.replacement)
		if out != v.want {
			t.Errorf("ToValidUTF8(%q, %q) = %q; want %q", v.input.s, v.input.replacement, out, v.want)
		}
	}
}

func TestDatabaseSql(t *testing.T) {
	ni32 := &sql.NullInt32{}
	v, err := ni32.Value()
	if err != nil {
		t.Errorf("NullInt32 Value fail. err: %v", err)
	}
	if v != nil {
		t.Errorf("NullInt32 Value not null")
	}

	nt := sql.NullTime{}
	v, err = nt.Value()
	if err != nil {
		t.Errorf("NullTime Value fail. err: %v", err)
	}
	if v != nil {
		t.Errorf("NullTime Value not null")
	}
}

func TestLog_Writer(t *testing.T) {
	switch x := log.Writer().(type) {
	case *os.File:
		if x != os.Stderr {
			t.Errorf("log.Writer() is not os.Stderr")
		}
	default:
		t.Errorf("log.Writer() is %T\n", log.Writer())
	}
}

//内联Once.Do Mutex/RWMutex
//For the uncontended cases on amd64, these changes make Once.Do twice as fast, and the Mutex/RWMutex methods up to 10% faster.
//Large Pool no longer increase stop-the-world pause times.
func TestSync(t *testing.T) {

}

func TestOS(t *testing.T) {
	fmt.Println(os.UserHomeDir())

	f, err := ioutil.TempFile("", "_Go_ErrIsExist")
	if err != nil {
		t.Fatalf("open ErrIsExist tempfile: %v", err)
	}
	name := f.Name()
	f.Close()
	os.Remove(name)

	f, err = os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		t.Fatalf("OpenFile fail. err: %v", err)
	}
	defer f.Close()

	_, err = f.WriteAt([]byte("xxx"), 0)
	if err == nil {
		t.Fatalf("WriteAt should have failed")
	}
}

func TestReflect(t *testing.T) {
	var testNil *int = nil
	tests := []struct {
		input interface{}
		want  bool
	}{
		{
			input: 0,
			want:  true,
		},
		{
			input: 0.0,
			want:  true,
		},
		{
			input: "",
			want:  true,
		},
		{
			input: false,
			want:  true,
		},
		{
			input: testNil,
			want:  true,
		},
		{
			input: new(chan struct{}),
			want:  false,
		},
		//{ //nil会导致panic
		//	input: nil,
		//	want:  true,
		//},
	}

	for _, v := range tests {
		out := reflect.ValueOf(v.input).IsZero()
		fmt.Println(reflect.ValueOf(v.input).Kind().String())
		if out != v.want {
			t.Fatalf("input: %v, want: %v, out: %v", v.input, v.want, out)
		}
	}

	swap := func(in []reflect.Value) []reflect.Value {
		return []reflect.Value{in[1], in[0]}
	}

	// makeSwap expects fptr to be a pointer to a nil function.
	// It sets that pointer to a new function created with MakeFunc.
	// When the function is invoked, reflect turns the arguments
	// into Values, calls swap, and then turns swap's result slice
	// into the values returned by the new function.
	makeSwap := func(fptr interface{}) {
		// fptr is a pointer to a function.
		// Obtain the function value itself (likely nil) as a reflect.Value
		// so that we can query its type and then set the value.
		fn := reflect.ValueOf(fptr).Elem()

		// Make a function of the right type.
		v := reflect.MakeFunc(fn.Type(), swap)

		// Assign it to the value fn represents.
		fn.Set(v)
	}

	// Make and call a swap function for ints.
	var intSwap func(int, int) (int, int)
	makeSwap(&intSwap)
	fmt.Println(intSwap(0, 1))

	// Make and call a swap function for float64s.
	var floatSwap func(float64, float64) (float64, float64)
	makeSwap(&floatSwap)
	fmt.Println(floatSwap(2.72, 3.14))
}

func TestReflectPanic(t *testing.T) {
	tests := []struct {
		input interface{}
	}{

		{
			input: nil,
		},
	}

	for _, v := range tests {
		var out error
		f := func(v interface{}) {
			defer func() {
				if err := recover(); err != nil {
					out = err.(error)
				}
			}()
			reflect.ValueOf(v).IsZero()
		}
		f(v.input)
		if out == nil {
			t.Fatalf("input: %v, out: %v", v.input, out)
		}
	}
}

func TestRuntime(t *testing.T) {
	for i := 0; i < 2; i++ {
		pc, file, line, ok := runtime.Caller(i)
		fmt.Println(runtime.FuncForPC(pc).Name(), file, line, ok)
	}
}

func TestTime(t *testing.T) {
	testTimes := []struct {
		input time.Time
		want  string
	}{
		{
			input: time.Date(2020, time.January, 30, 1, 1, 1, 1, time.Local),
			want:  "030",
		},
		{
			input: time.Date(2020, time.January, 31, 1, 1, 1, 1, time.Local),
			want:  "031",
		},
	}

	for _, v := range testTimes {
		out := v.input.Format("002")
		if out != v.want {
			t.Fatalf("input: %v, want: %v, out: %v", v.input, v.want, out)
		}
	}

	testDurations := []struct {
		input time.Duration
		want  int64
	}{
		{
			input: 2 * time.Second,
			want:  2_000,
		},
		{
			input: 1 * time.Second,
			want:  1_000,
		},
	}
	for _, v := range testDurations {
		out := v.input.Milliseconds()
		if out != v.want {
			t.Fatalf("input: %v, want: %v, out: %v", v.input, v.want, out)
		}
	}

	testDurations = []struct {
		input time.Duration
		want  int64
	}{
		{
			input: 2 * time.Second,
			want:  2_000_000,
		},
		{
			input: 1 * time.Second,
			want:  1_000_000,
		},
	}
	for _, v := range testDurations {
		out := v.input.Microseconds()
		if out != v.want {
			t.Fatalf("input: %v, want: %v, out: %v", v.input, v.want, out)
		}
	}
}
