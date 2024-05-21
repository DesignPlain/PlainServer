package types

type Composer_getEnvironmentConfigMasterAuthorizedNetworksConfigCidrBlock struct {
	// cidr_block must be specified in CIDR notation.
	CidrBlock string `json:"cidrBlock,omitempty" yaml:"cidrBlock,omitempty"`

	// display_name is a field for users to identify CIDR blocks.
	DisplayName string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
}
