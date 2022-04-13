package httparg

import "github.com/Bofry/httparg/internal"

type Processor struct {
	arg *internal.HttpArg

	hasError     bool
	errorHandler ErrorHandler
}

func NewProcessor(target interface{}, option Option) *Processor {
	var (
		arg          = internal.NewHttpArg(target)
		errorHandler = option.ErrorHandler
	)

	instance := Processor{
		arg:          arg,
		errorHandler: errorHandler,
	}

	return &instance
}

func (p *Processor) ProcessQueryString(query string) *Processor {
	if p.hasError {
		return p
	}

	var err error
	defer func() {
		if err != nil {
			p.throwError(err)
		}
	}()

	err = p.arg.ProcessQueryString(query)
	return p
}

func (p *Processor) ProcessContent(content []byte, contentType string) *Processor {
	if p.hasError {
		return p
	}

	var err error
	defer func() {
		if err != nil {
			p.throwError(err)
		}
	}()

	err = p.arg.ProcessContent(content, contentType)
	return p
}

func (p *Processor) Process(content []byte, processor ContentProcessor) *Processor {
	if p.hasError {
		return p
	}

	var err error
	defer func() {
		if err != nil {
			p.throwError(err)
		}
	}()

	err = p.arg.Process(content, processor)
	return p
}

func (p *Processor) Validate() {
	if p.hasError {
		return
	}

	var err error
	defer func() {
		if err != nil {
			p.throwError(err)
		}
	}()

	err = p.arg.Validate()
}

func (p *Processor) throwError(err error) {
	// set the hasError flag
	p.hasError = true

	errHandler := p.errorHandler
	if errHandler == nil {
		errHandler = RegistryService.CurrentErrorHandler()
	}
	if errHandler != nil {
		errHandler(err)
	}
}
