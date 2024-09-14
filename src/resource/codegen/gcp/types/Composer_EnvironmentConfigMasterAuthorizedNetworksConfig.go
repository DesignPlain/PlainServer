package types

type Composer_EnvironmentConfigMasterAuthorizedNetworksConfig struct {
	// cidr_blocks define up to 50 external networks that could access Kubernetes master through HTTPS.
	CidrBlocks []Composer_EnvironmentConfigMasterAuthorizedNetworksConfigCidrBlock `json:"cidrBlocks,omitempty" yaml:"cidrBlocks,omitempty"`

	// Whether or not master authorized networks is enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
