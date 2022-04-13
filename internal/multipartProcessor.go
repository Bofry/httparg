package internal

import (
	"bytes"
	"mime/multipart"

	"github.com/Bofry/structproto"
)

type MultipartProcessor struct {
	contentProcessService *ContentProcessService
	mediaType             string
	mediaParams           map[string]string
}

func (p *MultipartProcessor) Process(content []byte, target interface{}) error {
	var (
		reader = bytes.NewReader(content)
	)
	mr := multipart.NewReader(reader, p.mediaParams["boundary"])
	return p.bind(target, mr)
}

func (p *MultipartProcessor) bind(target interface{}, reader *multipart.Reader) error {
	var (
		provider = &MultipartContentBinder{
			reader:                reader,
			contentProcessService: p.contentProcessService,
		}
	)

	prototype, err := structproto.Prototypify(target,
		&structproto.StructProtoResolveOption{
			TagName: MultipartTagName,
		})

	if err != nil {
		return err
	}

	return prototype.Bind(provider)
}
