package types

type Compute_getInstanceTemplateShieldedInstanceConfig struct {
	// - Use a virtualized trusted platform module, which is a specialized computer chip you can use to encrypt objects like keys and certificates. Defaults to true.
	EnableVtpm bool `json:"enableVtpm,omitempty" yaml:"enableVtpm,omitempty"`

	// - Compare the most recent boot measurements to the integrity policy baseline and return a pair of pass/fail results depending on whether they match or not. Defaults to true.
	EnableIntegrityMonitoring bool `json:"enableIntegrityMonitoring,omitempty" yaml:"enableIntegrityMonitoring,omitempty"`

	// - Verify the digital signature of all boot components, and halt the boot process if signature verification fails. Defaults to false.
	EnableSecureBoot bool `json:"enableSecureBoot,omitempty" yaml:"enableSecureBoot,omitempty"`
}
