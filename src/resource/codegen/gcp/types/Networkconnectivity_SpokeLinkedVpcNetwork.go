package types

type Networkconnectivity_SpokeLinkedVpcNetwork struct {
	// IP ranges encompassing the subnets to be excluded from peering.
	ExcludeExportRanges []string `json:"excludeExportRanges,omitempty" yaml:"excludeExportRanges,omitempty"`

	// The URI of the VPC network resource.
	Uri string `json:"uri,omitempty" yaml:"uri,omitempty"`
}
