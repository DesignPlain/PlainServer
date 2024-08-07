package types

type Dataproc_ClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyInstanceSelectionResult struct {
	/*
	   The name of a Google Compute Engine machine type
	   to create for the node group. If not specified, GCP will default to a predetermined
	   computed value (currently `n1-standard-4`).
	*/
	MachineType string `json:"machineType,omitempty" yaml:"machineType,omitempty"`

	// Number of VM provisioned with the machine_type.
	VmCount int `json:"vmCount,omitempty" yaml:"vmCount,omitempty"`
}
