package reg

type srvDemo struct {
	baseMicroServiceMeta
}

func NewSrvDemo() MicroServiceMeta {
	return &srvDemo{baseMicroServiceMeta{
		name:     getServiceName("srv", "demo"),
		version:  "latest",
		metadata: make(map[string]string),
	}}
}
