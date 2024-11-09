package errorx

import (
	"fmt"
	"runtime"
)

type Error interface {
	error
	fmt.Formatter
	Unwrap() error
	Cause() error
	Code() int
	Type() ErrType
	Stack() Stack
}

type ErrType string

type Frame struct {
	Name string
	File string
	Line int
}

type Stack []Frame

// newStack 捕获当前的堆栈
func newStack() Stack {
	const depth = 32
	pcs := make([]uintptr, depth)
	n := runtime.Callers(3, pcs) // 跳过调用栈中的前几层

	stack := make(Stack, n)
	for i := 0; i < n; i++ {
		fn := runtime.FuncForPC(pcs[i])
		file, line := fn.FileLine(pcs[i])
		stack[i] = Frame{
			Name: fn.Name(),
			File: file,
			Line: line,
		}
	}
	return stack
}

// 定义 Error 接口实现
type customError struct {
	msg     string
	code    int
	errType ErrType
	cause   error
	stack   Stack
}

var _ Error = &customError{}

func (e *customError) Error() string {
	return e.msg
}

func (e *customError) Format(f fmt.State, c rune) {
	if len(e.errType) == 0 {
		fmt.Fprintf(f, "%s (code: %d)", e.msg, e.code)
	} else {
		fmt.Fprintf(f, "%s (code: %d, type: %s)", e.msg, e.code, e.errType)
	}
	if c == 'v' && f.Flag('+') {
		for _, frame := range e.stack {
			fmt.Fprintf(f, "\n\tat %s:%d in %s", frame.File, frame.Line, frame.Name)
		}
	}
}

func (e *customError) Unwrap() error {
	return e.cause
}

func (e *customError) Cause() error {
	return e.cause
}

func (e *customError) Code() int {
	return e.code
}

func (e *customError) Type() ErrType {
	return e.errType
}

func (e *customError) Stack() Stack {
	return e.stack
}

// Wrap 函数，用于包装已有的错误并添加栈信息
func Wrap(err error) Error {
	if err == nil {
		return nil
	}
	return &customError{
		msg:   err.Error(),
		cause: err,
		stack: newStack(),
	}
}

// New 函数，创建一个新的错误
func New(msg string) Error {
	return &customError{
		msg:     msg,
		code:    CODE_UNDEFINED, // 默认 code 为 CODE_UNDEFINED
		stack:   newStack(),
		errType: codeToErrTypeMap[CODE_UNDEFINED],
	}
}

// C 函数，创建带有指定 code 和 msg 的新错误
func C(code int, msg string) Error {
	errType := ERRTYPE_UNDEFINED
	if errTypeInfo, ok := codeToErrTypeMap[code]; ok {
		errType = errTypeInfo
	}
	return &customError{
		msg:     msg,
		code:    code,
		errType: errType,
		stack:   newStack(),
	}
}

// Cf 函数，创建带有指定 code 和格式化消息的错误
func Cf(code int, format string, args ...interface{}) Error {
	errType := ERRTYPE_UNDEFINED
	if errTypeInfo, ok := codeToErrTypeMap[code]; ok {
		errType = errTypeInfo
	}
	return &customError{
		msg:     fmt.Sprintf(format, args...),
		code:    code,
		errType: errType,
		stack:   newStack(),
	}
}
