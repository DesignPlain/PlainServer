package types

type Eks_ClusterAccessConfig struct {
	// The authentication mode for the cluster. Valid values are `CONFIG_MAP`, `API` or `API_AND_CONFIG_MAP`
	AuthenticationMode string `json:"authenticationMode,omitempty" yaml:"authenticationMode,omitempty"`

	// Whether or not to bootstrap the access config values to the cluster. Default is `true`.
	BootstrapClusterCreatorAdminPermissions bool `json:"bootstrapClusterCreatorAdminPermissions,omitempty" yaml:"bootstrapClusterCreatorAdminPermissions,omitempty"`
}
