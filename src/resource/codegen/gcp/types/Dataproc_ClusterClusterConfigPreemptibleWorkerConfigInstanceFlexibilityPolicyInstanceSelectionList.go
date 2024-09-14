package types

type Dataproc_ClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyInstanceSelectionList struct {
	// Full machine-type names, e.g. `"n1-standard-16"`.
	MachineTypes []string `json:"machineTypes,omitempty" yaml:"machineTypes,omitempty"`

	/*
	   Preference of this instance selection. A lower number means higher preference. Dataproc will first try to create a VM based on the machine-type with priority rank and fallback to next rank based on availability. Machine types and instance selections with the same priority have the same preference.

	   - - -
	*/
	Rank int `json:"rank,omitempty" yaml:"rank,omitempty"`
}
