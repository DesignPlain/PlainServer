package types

type Dataproc_ClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicy struct {
	// List of instance selection options that the group will use when creating new VMs.
	InstanceSelectionLists []Dataproc_ClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyInstanceSelectionList `json:"instanceSelectionLists,omitempty" yaml:"instanceSelectionLists,omitempty"`

	// A list of instance selection results in the group.
	InstanceSelectionResults []Dataproc_ClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicyInstanceSelectionResult `json:"instanceSelectionResults,omitempty" yaml:"instanceSelectionResults,omitempty"`
}
