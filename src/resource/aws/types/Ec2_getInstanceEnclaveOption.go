package types

type Ec2_getInstanceEnclaveOption struct {
	// Whether Nitro Enclaves are enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
