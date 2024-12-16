package app_error

import (
	"errors"
	"runtime"
	"strconv"
	"strings"
)

const delimiter = "-----"

type AppError struct {
	error      string
	stackTrace string
}

func New(error error) *AppError {
	err := new(AppError)
	if errors.As(error, &err) {
		return err
	}

	err.error = error.Error()
	err.stackTrace = err.getStackTrace()

	return err
}

func (e *AppError) Error() string {
	return e.error
}

func (e *AppError) StackTrace() string {
	return e.stackTrace
}

func (e *AppError) getStackTrace() string {
	var stackTrace strings.Builder
	stackTrace.WriteString("\n")
	stackTrace.WriteString(delimiter)

	frames := e.getRuntimeFrames()
	for {
		frame, hasNext := frames.Next()
		stackTrace.WriteString(e.buildStackTraceLine(frame))

		if hasNext {
			stackTrace.WriteString("\n" + delimiter + "\n")
		} else {
			break
		}
	}

	return stackTrace.String()
}

func (e *AppError) buildStackTraceLine(frame runtime.Frame) string {
	var stackTraceLine strings.Builder
	stackTraceLine.WriteString("\"File: " + frame.File + "\",\n")
	stackTraceLine.WriteString("\"Func: " + frame.Func.Name() + "\",\n")
	stackTraceLine.WriteString("\"Line: " + strconv.Itoa(frame.Line) + "\"")

	return stackTraceLine.String()
}

func (e *AppError) getRuntimeFrames() *runtime.Frames {
	stack := make([]uintptr, 10)
	n := runtime.Callers(4, stack[:])
	return runtime.CallersFrames(stack[:n])
}
