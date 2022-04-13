package internal

type HttpArg struct {
	target interface{}

	contentProcessService *ContentProcessService
}

func NewHttpArg(target interface{}) *HttpArg {
	instance := HttpArg{
		target:                target,
		contentProcessService: ContentProcessServiceInstance,
	}

	return &instance
}

func (arg *HttpArg) ProcessQueryString(query string) error {
	return arg.contentProcessService.ProcessQueryArgs(arg.target, query)
}

func (arg *HttpArg) ProcessContent(content []byte, contentType string) error {
	return arg.contentProcessService.Process(arg.target, content, contentType)
}

func (arg *HttpArg) Process(content []byte, processor ContentProcessor) error {
	return processor(content, arg.target)
}

func (arg *HttpArg) Validate() error {
	v, ok := arg.target.(Validatable)
	if ok {
		return v.Validate()
	}
	return nil
}
