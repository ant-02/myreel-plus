package errno

import (
	"fmt"
	"io"
	"runtime"
	"strings"
)

type ErrNo struct {
	ErrorMsg string
	err      error
	stack    *stack
}

func (e *ErrNo) Error() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("%s", e.ErrorMsg)
}

func (e *ErrNo) WithStack() error {
	e.stack = callers(3)
	return e
}

func (e *ErrNo) Unwrap() error {
	return e.err
}

func (e *ErrNo) StackTrace() string {
	if e == nil {
		return ""
	}

	var builder strings.Builder
	builder.WriteString(e.Error())
	builder.WriteString("\n")

	if e.stack != nil && len(*e.stack) > 0 {
		frames := runtime.CallersFrames(*e.stack)
		for {
			frame, more := frames.Next()
			if frame.File == "" {
				if !more {
					break
				}
				continue // 跳过无效帧
			}
			builder.WriteString(fmt.Sprintf("\t%s:%d %s\n", frame.File, frame.Line, frame.Function))
			if !more {
				break
			}
		}
	} else if e.err != nil {
		if next, ok := e.err.(*ErrNo); ok {
			builder.WriteString(next.StackTrace())
		} else {
			builder.WriteString(e.err.Error())
			builder.WriteString("\n")
		}
	}
	return builder.String()
}

func (e *ErrNo) Format(st fmt.State, verb rune) {
	switch verb {
	case 'v':
		if st.Flag('+') {
			io.WriteString(st, e.StackTrace())
			return
		}
		fallthrough
	case 's':
		io.WriteString(st, e.Error())
	}
}

func Errorf(err error, msg string, args ...interface{}) error {
	return &ErrNo{
		ErrorMsg: fmt.Sprintf(msg, args...),
		err:      err,
		stack:    nil,
	}
}
