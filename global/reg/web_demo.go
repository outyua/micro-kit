package reg

type webDemo struct {
	baseMicroServiceMeta
}

func NewWebDemo() MicroServiceMeta {
	return &webDemo{baseMicroServiceMeta{
		name:     getServiceName("web", "demo"),
		version:  "latest",
		metadata: make(map[string]string),
	}}
}
