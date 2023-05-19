package internal

import (
	"github.com/Bofry/structproto/reflecting"
)

const (
	HEADER_CONTENT_DISPOSITION = "Content-Disposition"
	HEADER_CONTENT_TYPE        = "Content-Type"

	True = "true"

	JsonTagName      = "json"
	MultipartTagName = "multipart"
)

var (
	ContentProcessServiceInstance = new(ContentProcessService)
	ContentProcessRegistryService = newContentProcessRegistry(ContentProcessServiceInstance)

	assignZero = reflecting.AssignZero
)

type (
	Validatable interface {
		Validate() error
	}
)

type (
	StringContentProcessor func(content string, target interface{}) error
	ContentProcessor       func(content []byte, target interface{}) error
)
