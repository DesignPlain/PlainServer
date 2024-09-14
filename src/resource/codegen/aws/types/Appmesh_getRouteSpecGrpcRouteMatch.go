package types

type Appmesh_getRouteSpecGrpcRouteMatch struct {
	//
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	//
	Prefix string `json:"prefix,omitempty" yaml:"prefix,omitempty"`

	//
	ServiceName string `json:"serviceName,omitempty" yaml:"serviceName,omitempty"`

	//
	Metadatas []Appmesh_getRouteSpecGrpcRouteMatchMetadata `json:"metadatas,omitempty" yaml:"metadatas,omitempty"`

	//
	MethodName string `json:"methodName,omitempty" yaml:"methodName,omitempty"`
}
