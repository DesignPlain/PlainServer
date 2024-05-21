package types

type Dataproc_ClusterClusterConfigPreemptibleWorkerConfig struct {
	// Disk Config
	DiskConfig Dataproc_ClusterClusterConfigPreemptibleWorkerConfigDiskConfig `json:"diskConfig,omitempty" yaml:"diskConfig,omitempty"`

	// Instance flexibility Policy allowing a mixture of VM shapes and provisioning models.
	InstanceFlexibilityPolicy Dataproc_ClusterClusterConfigPreemptibleWorkerConfigInstanceFlexibilityPolicy `json:"instanceFlexibilityPolicy,omitempty" yaml:"instanceFlexibilityPolicy,omitempty"`

	// List of preemptible instance names which have been assigned to the cluster.
	InstanceNames []string `json:"instanceNames,omitempty" yaml:"instanceNames,omitempty"`

	/*
	   Specifies the number of preemptible nodes to create.
	   Defaults to 0.
	*/
	NumInstances int `json:"numInstances,omitempty" yaml:"numInstances,omitempty"`

	/*
	   Specifies the preemptibility of the secondary workers. The default value is `PREEMPTIBLE`
	   Accepted values are:
	   - PREEMPTIBILITY_UNSPECIFIED
	   - NON_PREEMPTIBLE
	   - PREEMPTIBLE
	*/
	Preemptibility string `json:"preemptibility,omitempty" yaml:"preemptibility,omitempty"`
}
