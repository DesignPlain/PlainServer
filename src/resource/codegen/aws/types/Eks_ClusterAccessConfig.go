package types

type Eks_ClusterAccessConfig struct {
	// Whether or not to bootstrap the access config values to the cluster. Default is `false`.
	BootstrapClusterCreatorAdminPermissions bool `json:"bootstrapClusterCreatorAdminPermissions,omitempty" yaml:"bootstrapClusterCreatorAdminPermissions,omitempty"`

	// The authentication mode for the cluster. Valid values are `CONFIG_MAP`, `API` or `API_AND_CONFIG_MAP`
	AuthenticationMode string `json:"authenticationMode,omitempty" yaml:"authenticationMode,omitempty"`
}
