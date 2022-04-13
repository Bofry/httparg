package httparg

import (
	"github.com/Bofry/httparg/internal"
)

var RegistryService Registry

type Registry struct{}

func (r Registry) SetupErrorHandler(fn ErrorHandler) {
	if errorHandler != nil {
		logger.Panic("cannot setup global ErrorHandleFunc")
	}
	initErrorHandler(fn)
}

func (r Registry) CurrentErrorHandler() ErrorHandler {
	if errorHandler != nil {
		return errorHandler
	}
	return stdErrorHandler
}

func (r Registry) RegisterContentProcessor(mediatype string, processor ContentProcessor) bool {
	return internal.ContentProcessRegistryService.RegisterContentProcessor(mediatype, processor)
}
