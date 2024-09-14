package types

type Dataproc_WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig struct {
	// Defines whether instances have [Integrity Monitoring](https://cloud.google.com/compute/shielded-vm/docs/shielded-vm#integrity-monitoring) enabled.
	EnableIntegrityMonitoring bool `json:"enableIntegrityMonitoring,omitempty" yaml:"enableIntegrityMonitoring,omitempty"`

	// Defines whether instances have [Secure Boot](https://cloud.google.com/compute/shielded-vm/docs/shielded-vm#secure-boot) enabled.
	EnableSecureBoot bool `json:"enableSecureBoot,omitempty" yaml:"enableSecureBoot,omitempty"`

	// Defines whether instances have the [vTPM](https://cloud.google.com/compute/shielded-vm/docs/shielded-vm#vtpm) enabled.
	EnableVtpm bool `json:"enableVtpm,omitempty" yaml:"enableVtpm,omitempty"`
}
