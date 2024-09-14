package types

type Workstations_WorkstationConfigHostGceInstanceShieldedInstanceConfig struct {
	// Whether the instance has integrity monitoring enabled.
	EnableIntegrityMonitoring bool `json:"enableIntegrityMonitoring,omitempty" yaml:"enableIntegrityMonitoring,omitempty"`

	// Whether the instance has Secure Boot enabled.
	EnableSecureBoot bool `json:"enableSecureBoot,omitempty" yaml:"enableSecureBoot,omitempty"`

	// Whether the instance has the vTPM enabled.
	EnableVtpm bool `json:"enableVtpm,omitempty" yaml:"enableVtpm,omitempty"`
}
