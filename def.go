package httparg

import (
	"log"
	"os"
	"sync"

	"github.com/Bofry/arg"
	"github.com/Bofry/httparg/form"
	"github.com/Bofry/httparg/internal"
	"github.com/Bofry/httparg/json"
	"github.com/Bofry/httparg/querystring"
)

const (
	LOGGER_PREFIX string = "[httparg] "
)

var (
	logger *log.Logger = log.New(os.Stdout, LOGGER_PREFIX, log.LstdFlags|log.Lmsgprefix)

	errorHandlerOnce sync.Once
	errorHandler     ErrorHandler

	stdErrorHandler = func(err error) { panic(err) }

	canonicalContentProcessors = map[string]ContentProcessor{
		"application/octet-stream":          nil,
		"text/plain":                        nil,
		"application/x-www-form-urlencoded": form.Process,
		"application/json":                  json.Process,
	}
)

// interface
type (
	InvalidArgumentError = arg.InvalidArgumentError

	Validatable interface {
		Validate() error
	}
)

// function
type (
	ContentProcessor = internal.ContentProcessor
	ErrorHandler     func(err error)
)

func init() {
	internal.ContentProcessRegistryService.Setup(
		querystring.Process,
		canonicalContentProcessors,
	)
}

func initErrorHandler(fn ErrorHandler) {
	errorHandlerOnce.Do(func() {
		if fn != nil {
			errorHandler = fn
		}
	})
}
