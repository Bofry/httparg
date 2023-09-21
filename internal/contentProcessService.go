package internal

import (
	"fmt"
	"strings"
)

type ContentProcessService struct {
	queryArgsProcessor StringContentProcessor

	processors map[string]ContentProcessor
}

func (service *ContentProcessService) Process(target interface{}, content []byte, contentType string) error {
	var (
		mediatype ContentType = ContentType{}
		err       error
	)

	// remove all spaces
	contentType = strings.Trim(contentType, " ")

	if contentType != "" {
		mediatype, err = ParseContentType(contentType)
		if err != nil {
			return err
		}
	}

	processor := service.getProcessor(mediatype)
	if processor == nil {
		return fmt.Errorf("cannot process specified content-type '%s'", mediatype)
	}
	return processor(content, target)
}

func (service *ContentProcessService) ProcessQueryArgs(target interface{}, content string) error {
	var (
		proc = service.queryArgsProcessor
	)
	if proc == nil {
		return fmt.Errorf("cannot process query string")
	}
	return proc(content, target)
}

func (service *ContentProcessService) getProcessor(contentType ContentType) ContentProcessor {
	if service.processors == nil {
		return nil
	}

	switch {
	case contentType.IsMultipartTypes():
		processor := &MultipartProcessor{
			contentProcessService: service,
			mediaType:             contentType.mediatype,
			mediaParams:           contentType.params,
		}
		return processor.Process

	default:
		if v, ok := service.processors[contentType.mediatype]; ok {
			return v
		}
		return nil
	}
}
