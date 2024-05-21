package types

type Composer_getEnvironmentConfigMasterAuthorizedNetworksConfig struct {
	// cidr_blocks define up to 50 external networks that could access Kubernetes master through HTTPS.
	CidrBlocks []Composer_getEnvironmentConfigMasterAuthorizedNetworksConfigCidrBlock `json:"cidrBlocks,omitempty" yaml:"cidrBlocks,omitempty"`

	// Whether or not master authorized networks is enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
