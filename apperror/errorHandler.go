package apperror

import (
	"fmt"

	"github.com/miyamoto-jo/go-batch-sample/logger"
)

const (
	ExitCodeOK    = 0
	ExitCodeError = 1
)

type ExitError struct {
	exitCode int
	err      error
}

func (e *ExitError) Error() string {
	if e.err == nil {
		return ""
	}
	return fmt.Sprintf("%v", e.err)
}

func NewExitError(code int, err error) *ExitError {
	return &ExitError{
		exitCode: code,
		err:      err,
	}
}

func ExitHandler(err error) int {
	if err == nil {
		return ExitCodeOK
	}

	if e, ok := err.(*ExitError); ok {
		logger.Error("err: ", err)
		return e.exitCode
	}

	if e, ok := err.(error); ok {
		logger.Error("err: ", e)
		return ExitCodeError
	}
	return ExitCodeOK
}
