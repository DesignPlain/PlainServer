package types

type Container_NodePoolNodeConfigShieldedInstanceConfig struct {
	// Defines whether the instance has Secure Boot enabled.
	EnableSecureBoot bool `json:"enableSecureBoot,omitempty" yaml:"enableSecureBoot,omitempty"`

	// Defines whether the instance has integrity monitoring enabled.
	EnableIntegrityMonitoring bool `json:"enableIntegrityMonitoring,omitempty" yaml:"enableIntegrityMonitoring,omitempty"`
}
