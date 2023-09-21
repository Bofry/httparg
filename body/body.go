package body

import (
	"github.com/Bofry/httparg/internal"
	"github.com/Bofry/structproto"
)

const (
	TagName = "body"
)

var _ internal.ContentProcessor = Process

func Process(content []byte, target interface{}) error {
	provider := internal.NewBodyBinder(content)

	prototype, err := structproto.Prototypify(target,
		&structproto.StructProtoResolveOption{
			TagName: TagName,
		})
	if err != nil {
		return err
	}

	return prototype.Bind(provider)
}
