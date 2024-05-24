package types

type Eks_getClusterAccessConfig struct {
	// Values returned are `CONFIG_MAP`, `API` or `API_AND_CONFIG_MAP`
	AuthenticationMode string `json:"authenticationMode,omitempty" yaml:"authenticationMode,omitempty"`
}
