package internal

import (
	"io"
	"io/ioutil"
	"mime/multipart"
	"reflect"

	"github.com/Bofry/structproto"
	"github.com/Bofry/structproto/valuebinder"
)

var _ structproto.StructBinder = new(MultipartContentBinder)

type MultipartContentBinder struct {
	reader                *multipart.Reader
	contentProcessService *ContentProcessService

	fieldnames []string
}

func (binder *MultipartContentBinder) Init(context *structproto.StructProtoContext) error {
	for {
		part, err := binder.reader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		var name = part.FormName()
		if len(name) > 0 {
			if rv, ok := context.Field(name); ok {
				var (
					contentType = part.Header.Get(HEADER_CONTENT_TYPE)
				)

				body, err := ioutil.ReadAll(part)
				if err != nil {
					return err
				}

				if len(contentType) > 0 {
					rv = assignZero(rv)
					err = binder.contentProcessService.Process(rv, body, contentType)
				} else {
					err = valuebinder.BuildBytesBinder(rv).Bind(body)
				}

				if err != nil {
					return err
				}

				binder.fieldnames = append(binder.fieldnames, name)
			}
		}
	}

	return nil
}

func (binder *MultipartContentBinder) Bind(field structproto.FieldInfo, rv reflect.Value) error {
	// ignore
	return nil
}

func (binder *MultipartContentBinder) Deinit(context *structproto.StructProtoContext) error {
	return context.CheckIfMissingRequiredFields(func() <-chan string {
		c := make(chan string, 1)
		go func() {
			for _, v := range binder.fieldnames {
				c <- v
			}
			close(c)
		}()
		return c
	})
}
