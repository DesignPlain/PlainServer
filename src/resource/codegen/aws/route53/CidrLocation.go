package route53

type CidrLocation struct {
	// CIDR blocks for the location.
	CidrBlocks []string `json:"cidrBlocks,omitempty" yaml:"cidrBlocks,omitempty"`

	// The ID of the CIDR collection to update.
	CidrCollectionId string `json:"cidrCollectionId,omitempty" yaml:"cidrCollectionId,omitempty"`

	// Name for the CIDR location.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
