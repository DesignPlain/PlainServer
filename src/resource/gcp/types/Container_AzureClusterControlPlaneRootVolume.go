package types

type Container_AzureClusterControlPlaneRootVolume struct {
	// Optional. The size of the disk, in GiBs. When unspecified, a default value is provided. See the specific reference in the parent resource.
	SizeGib int `json:"sizeGib,omitempty" yaml:"sizeGib,omitempty"`
}
