package types

type Appmesh_VirtualNodeSpecServiceDiscovery struct {
	// Any AWS Cloud Map information for the virtual node.
	AwsCloudMap Appmesh_VirtualNodeSpecServiceDiscoveryAwsCloudMap `json:"awsCloudMap,omitempty" yaml:"awsCloudMap,omitempty"`

	// DNS service name for the virtual node.
	Dns Appmesh_VirtualNodeSpecServiceDiscoveryDns `json:"dns,omitempty" yaml:"dns,omitempty"`
}
