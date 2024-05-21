package types

type Dataproc_ClusterClusterConfigGceClusterConfigShieldedInstanceConfig struct {
	/*
	   Defines whether instances have integrity monitoring enabled.

	   - - -
	*/
	EnableIntegrityMonitoring bool `json:"enableIntegrityMonitoring,omitempty" yaml:"enableIntegrityMonitoring,omitempty"`

	// Defines whether instances have Secure Boot enabled.
	EnableSecureBoot bool `json:"enableSecureBoot,omitempty" yaml:"enableSecureBoot,omitempty"`

	// Defines whether instances have the [vTPM](https://cloud.google.com/security/shielded-cloud/shielded-vm#vtpm) enabled.
	EnableVtpm bool `json:"enableVtpm,omitempty" yaml:"enableVtpm,omitempty"`
}
