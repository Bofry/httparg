package json

import (
	"github.com/Bofry/httparg/internal"
)

const (
	TagName = internal.JsonTagName
)

var _ internal.ContentProcessor = Process

func Process(content []byte, target interface{}) error {
	rv, err := internal.Indirect(target)
	if err != nil {
		return err
	}

	binder := internal.BuildJsonValueBinder(rv)
	err = binder.Bind(content)
	if err != nil {
		return err
	}
	return nil
}
