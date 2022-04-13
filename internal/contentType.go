package internal

import (
	"mime"
	"strings"
)

type ContentType struct {
	mediatype string
	params    map[string]string
}

func ParseContentType(contentType string) (ContentType, error) {
	var instance ContentType
	mediatype, params, err := mime.ParseMediaType(contentType)
	if err != nil {
		return instance, err
	}

	return ContentType{
		mediatype: mediatype,
		params:    params,
	}, nil
}

// RFC 7231 section 3.1.1.4 & RFC 2046 section 5.1.1
func (t ContentType) IsMultipartTypes() bool {
	return strings.HasPrefix(t.mediatype, "multipart/")
}
