package errors

import (
	"errors"
	"fmt"
)

var errTest = errors.New("test error")

// https://github.com/golang/go/wiki/ErrorValueFAQ
type unWrapError struct {
	err error
}

func (u *unWrapError) Error() string {
	return fmt.Sprintf("unWrapError: %v", u.err)
}

func (u *unWrapError) Unwrap() error {
	return u.err
}

//If you check for an error type using a type assertion or type switch, use errors.As instead. Example:
//  if e, ok := err.(*os.PathError); ok
//becomes
//	var e *os.PathError
//	if errors.As(err, &e)
type asError struct {
	unWrapError
}

func (a *asError) Error() string {
	return fmt.Sprintf("asError: %v", a.err)
}

func (a *asError) As(err interface{}) bool {
	fmt.Printf("As %T\n", err)
	switch x := err.(type) {
	case **errorT:
		*x = &errorT{
			unWrapError: unWrapError{
				err: a.err,
			},
		}
	default:
		return false
	}
	return true
}

type errorT struct {
	unWrapError
}

func (e *errorT) Error() string {
	return fmt.Sprintf("errorT: %v", e.err)
}

//You need to be prepared that errors you get may be wrapped.
//If you currently compare errors using ==, use errors.Is instead. Example:
//	if err == io.ErrUnexpectedEOF
//becomes
// 	if errors.Is(err, io.ErrUnexpectedEOF)
type isError struct {
	unWrapError
}

func (i *isError) Error() string {
	return fmt.Sprintf("isError: %v", i.err)
}

func (i *isError) Is(err error) bool {

	switch err.(type) {
	case *isError:
		return true
	}
	return false
}
