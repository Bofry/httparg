package internal

type ContentProcessRegistry struct {
	service *ContentProcessService
}

func newContentProcessRegistry(service *ContentProcessService) *ContentProcessRegistry {
	return &ContentProcessRegistry{
		service: service,
	}
}

func (registry *ContentProcessRegistry) Setup(
	queryArgsProcessor StringContentProcessor,
	processors map[string]ContentProcessor) {

	registry.service.queryArgsProcessor = queryArgsProcessor
	registry.service.processors = processors
}

func (registry *ContentProcessRegistry) Get(mediatype string) ContentProcessor {
	if registry.service.processors != nil {
		processor, ok := registry.service.processors[mediatype]
		if ok {
			return processor
		}
	}
	return nil
}

func (registry *ContentProcessRegistry) RegisterContentProcessor(mediatype string, processor ContentProcessor) bool {
	if registry.service.processors != nil {
		if processor == nil {
			delete(registry.service.processors, mediatype)
		} else {
			registry.service.processors[mediatype] = processor
		}
		return true
	}
	return false
}
