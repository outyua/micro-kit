package reg

import "fmt"

// namespace config

const NAMESPACE = "go.micro"

// 注册所有服务
type MicroServiceMeta interface {
	GetName() string
	GerVersion() string
	GetMetaData() map[string]string
}

func getServiceName(typ, name string) string {
	return fmt.Sprintf("%s.%s.%s", NAMESPACE, typ, name)
}

type baseMicroServiceMeta struct {
	name     string
	version  string
	metadata map[string]string
}

func (s *baseMicroServiceMeta) GetName() string {
	return s.name
}

func (s *baseMicroServiceMeta) GerVersion() string {
	return s.version
}

func (s *baseMicroServiceMeta) GetMetaData() map[string]string {
	return s.metadata
}
