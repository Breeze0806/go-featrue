package errors

import (
	"errors"
	"fmt"
	"testing"
)

func Test_Unwrap(t *testing.T) {
	testCases := []struct {
		input error
		want  error
	}{
		{
			input: &unWrapError{
				err: errTest,
			},
			want: errTest,
		},
		{
			input: errTest,
			want:  nil,
		},
	}

	for _, v := range testCases {
		out := errors.Unwrap(v.input)
		if !errors.Is(out, v.want) {
			t.Fatalf("out: %v want: %v input: %+v", out, v.want, v.input)
		}
	}
}

func Test_Is(t *testing.T) {
	testCases := []struct {
		input struct {
			err    error
			target error
		}
		want bool
	}{
		{
			input: struct {
				err    error
				target error
			}{
				err: &unWrapError{
					err: errTest,
				},
				target: errTest,
			},
			want: true,
		},
		{
			input: struct {
				err    error
				target error
			}{
				err: &isError{
					unWrapError: unWrapError{
						err: errTest,
					},
				},
				target: errTest,
			},
			want: true,
		},
		{
			input: struct {
				err    error
				target error
			}{
				err: &isError{
					unWrapError: unWrapError{
						err: errTest,
					},
				},
				target: &unWrapError{
					err: errTest,
				},
			},
			want: false,
		},
		{
			input: struct {
				err    error
				target error
			}{
				err: &isError{
					unWrapError: unWrapError{
						err: errTest,
					},
				},
				target: &isError{
					unWrapError: unWrapError{
						err: errTest,
					},
				},
			},
			want: false,
		},
	}

	for _, v := range testCases {
		out := errors.Is(v.input.err, v.input.target)
		if out != v.want {
			t.Fatalf("out: %v want: %v input: %#v", out, v.want, v.input)
		}
	}
}

func Test_As(t *testing.T) {
	testCases := []struct {
		input struct {
			err    error
			target interface{}
		}
		want bool
	}{
		{
			input: struct {
				err    error
				target interface{}
			}{
				err: &unWrapError{
					err: errTest,
				},
				target: new(*unWrapError),
			},
			want: true,
		},
		{
			input: struct {
				err    error
				target interface{}
			}{
				err: &asError{
					unWrapError: unWrapError{
						err: errTest,
					},
				},
				target: new(*asError),
			},
			want: true,
		},
		{
			input: struct {
				err    error
				target interface{}
			}{
				err: &unWrapError{
					err: errTest,
				},
				target: new(error),
			},
			want: true,
		},
		{
			input: struct {
				err    error
				target interface{}
			}{
				err: &unWrapError{
					err: errTest,
				},
				target: new(*errorT),
			},
			want: false,
		},
		{
			input: struct {
				err    error
				target interface{}
			}{
				err: &asError{
					unWrapError: unWrapError{
						err: errTest,
					},
				},
				target: new(*errorT),
			},
			want: true,
		},
	}

	for i, v := range testCases {
		fmt.Println("=======", i, "=========")
		out := errors.As(v.input.err, v.input.target)
		if out != v.want {
			t.Fatalf("out: %v want: %v input: %#v", out, v.want, v.input)
		}
		switch x := v.input.target.(type) {
		case **unWrapError:
			fmt.Println(*x)
		case **asError:
			fmt.Println(*x)
		case *error:
			fmt.Println(*x)
		case **errorT:
			fmt.Println(*x)
		}
	}
}

func Test_AsPanic(t *testing.T) {
	testCases := []struct {
		input struct {
			err    error
			target interface{}
		}
		want string
	}{
		{
			input: struct {
				err    error
				target interface{}
			}{
				err: &unWrapError{
					err: errTest,
				},
				target: nil,
			},
			want: "errors: target cannot be nil",
		},
		{
			input: struct {
				err    error
				target interface{}
			}{
				err: &asError{
					unWrapError: unWrapError{
						err: errTest,
					},
				},
				target: new(asError),
			},
			want: "errors: *target must be interface or implement error",
		},
	}

	for _, v := range testCases {
		out := ""
		f := func(v struct {
			err    error
			target interface{}
		}) {
			defer func() {
				if err := recover(); err != nil {
					out = err.(string)
				}
			}()
			errors.As(v.err, v.target)
		}
		f(v.input)
		if out != v.want {
			t.Fatalf("out: %v want: %v input: %#v", out, v.want, v.input)
		}
	}
}

func Test_Errorf(t *testing.T) {
	testCases := []struct {
		input error
		want  error
	}{
		{
			input: &unWrapError{
				err: errTest,
			},
			want: errTest,
		},
		{
			input: errTest,
			want:  nil,
		},
	}

	for i, v := range testCases {
		err := fmt.Errorf("xx: %w\n", v.input)

		if _, ok := err.(interface{ As(interface{}) bool }); !ok {
			fmt.Println(i, err)
		}
	}
}
