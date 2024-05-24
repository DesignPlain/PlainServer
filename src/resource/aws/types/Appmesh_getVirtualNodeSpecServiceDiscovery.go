package types

type Appmesh_getVirtualNodeSpecServiceDiscovery struct {
	//
	AwsCloudMaps []Appmesh_getVirtualNodeSpecServiceDiscoveryAwsCloudMap `json:"awsCloudMaps,omitempty" yaml:"awsCloudMaps,omitempty"`

	//
	Dns []Appmesh_getVirtualNodeSpecServiceDiscoveryDn `json:"dns,omitempty" yaml:"dns,omitempty"`
}
