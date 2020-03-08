package reg

type srvAdmin struct {
	baseMicroServiceMeta
}

func NewSrvAdmin() MicroServiceMeta {
	return &srvAdmin{baseMicroServiceMeta{
		name:     getServiceName("srv", "admin"),
		version:  "latest",
		metadata: make(map[string]string),
	}}
}
