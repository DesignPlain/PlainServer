package types

type Dataproc_ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfig struct {
	// The node pool configuration.
	Config Dataproc_ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigConfig `json:"config,omitempty" yaml:"config,omitempty"`

	/*
	   The list of Compute Engine zones where node pool nodes associated
	   with a Dataproc on GKE virtual cluster will be located.
	   - - -
	*/
	Locations []string `json:"locations,omitempty" yaml:"locations,omitempty"`

	/*
	   The autoscaler configuration for this node pool.
	   The autoscaler is enabled only when a valid configuration is present.
	*/
	Autoscaling Dataproc_ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfigAutoscaling `json:"autoscaling,omitempty" yaml:"autoscaling,omitempty"`
}
