package types

type Compute_InstanceFromMachineImageShieldedInstanceConfig struct {
	// Whether the instance uses vTPM.
	EnableVtpm bool `json:"enableVtpm,omitempty" yaml:"enableVtpm,omitempty"`

	// Whether integrity monitoring is enabled for the instance.
	EnableIntegrityMonitoring bool `json:"enableIntegrityMonitoring,omitempty" yaml:"enableIntegrityMonitoring,omitempty"`

	// Whether secure boot is enabled for the instance.
	EnableSecureBoot bool `json:"enableSecureBoot,omitempty" yaml:"enableSecureBoot,omitempty"`
}
