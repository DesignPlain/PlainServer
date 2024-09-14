package types

type Appmesh_getVirtualNodeSpecServiceDiscoveryAwsCloudMap struct {
	//
	ServiceName string `json:"serviceName,omitempty" yaml:"serviceName,omitempty"`

	//
	Attributes map[string]string `json:"attributes,omitempty" yaml:"attributes,omitempty"`

	//
	NamespaceName string `json:"namespaceName,omitempty" yaml:"namespaceName,omitempty"`
}
