package httparg

func Args(target interface{}) *Processor {
	return NewProcessor(target, Option{})
}

func ArgsWithOption(target interface{}, option Option) *Processor {
	return NewProcessor(target, option)
}
