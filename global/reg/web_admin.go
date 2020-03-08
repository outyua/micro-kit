package reg

type webAdmin struct {
	baseMicroServiceMeta
}

func NewWebAdmin() MicroServiceMeta {
	return &webAdmin{baseMicroServiceMeta{
		name:     getServiceName("web", "admin"),
		version:  "latest",
		metadata: make(map[string]string),
	}}
}
