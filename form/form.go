package form

import (
	"net/url"

	"github.com/Bofry/httparg/internal"
	"github.com/Bofry/structproto"
)

const (
	TagName = "form"
	True    = "true"
)

var _ internal.ContentProcessor = Process

func Process(content []byte, target interface{}) error {
	values, err := url.ParseQuery(string(content))
	if err != nil {
		return err
	}

	provider := internal.NewQueryArgsBinder(values)

	prototype, err := structproto.Prototypify(target,
		&structproto.StructProtoResolveOption{
			TagName: TagName,
		})
	if err != nil {
		return err
	}

	return prototype.Bind(provider)
}
