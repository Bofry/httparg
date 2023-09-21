package internal

import (
	"reflect"

	"github.com/Bofry/structproto"
	"github.com/Bofry/structproto/valuebinder"
)

var _ structproto.StructBinder = new(BodyBinder)

type BodyBinder struct {
	content []byte
}

func NewBodyBinder(content []byte) *BodyBinder {
	instance := &BodyBinder{
		content: content,
	}
	return instance
}

// Init implements structproto.StructBinder.
func (binder *BodyBinder) Init(context *structproto.StructProtoContext) error {
	return nil
}

// Bind implements structproto.StructBinder.
func (binder *BodyBinder) Bind(field structproto.FieldInfo, rv reflect.Value) error {
	return valuebinder.BuildBytesBinder(rv).Bind(binder.content)
}

// Deinit implements structproto.StructBinder.
func (binder *BodyBinder) Deinit(context *structproto.StructProtoContext) error {
	return nil
}
