package common

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
)

type Error struct {
	msg   string
	loc   string
	cause error
}

const (
	PrefixName = "github.com/FelixHolmes/glc/"
)

// WrapError 包装err
func WrapError(cause error) *Error {
	return newError("").WithCause(cause)
}

// Errorf 创建一个新的err
func Errorf(msg string, a ...interface{}) *Error {
	if len(a) > 0 {
		msg = fmt.Sprintf(msg, a...)
	}
	return newError(msg)
}

func fixLocPreStr(s string) string {
	t := strings.SplitN(s, "glc", 2)
	if len(t) == 2 {
		return warpPreStr(t[1])
	}
	return warpPreStr(s)
}

func warpPreStr(s string) string {
	var buffer bytes.Buffer
	buffer.WriteString(PrefixName)
	buffer.WriteString(s)
	return buffer.String()
}

func newError(msg string) *Error {
	pc, fn, ln, _ := runtime.Caller(2)
	fn = fixLocPreStr(fn)
	funcName := fixLocPreStr(runtime.FuncForPC(pc).Name())
	return &Error{
		msg: msg,
		loc: fmt.Sprintf("%s [%s:%v]", funcName, fn, ln),
	}
}

func (e *Error) Error() string {
	if e.cause == nil {
		return e.msg
	}
	if e.msg == "" {
		return e.cause.Error()
	}
	return e.msg + ": " + e.cause.Error()
}

func (e *Error) Stack() string {
	if c, ok := e.cause.(*Error); ok {
		return c.Stack() + "\n\t" + e.loc
	} else {
		return e.loc
	}
}

func (e *Error) WithCause(cause error) *Error {
	e.cause = cause
	return e
}
