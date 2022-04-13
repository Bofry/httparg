package querystring

import (
	"net/url"

	"github.com/Bofry/httparg/internal"
	"github.com/Bofry/structproto"
)

const (
	TagName = "query"
)

var _ internal.StringContentProcessor = Process

func Process(content string, target interface{}) error {
	values, err := url.ParseQuery(content)
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
