package types

type Dataproc_ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget struct {
	/*
	   Node group roles.
	   One of `"DRIVER"`.
	*/
	Roles []string `json:"roles,omitempty" yaml:"roles,omitempty"`

	// The target GKE node pool.
	NodePool string `json:"nodePool,omitempty" yaml:"nodePool,omitempty"`

	/*
	   The configuration for the GKE node pool.
	   If specified, Dataproc attempts to create a node pool with the specified shape.
	   If one with the same name already exists, it is verified against all specified fields.
	   If a field differs, the virtual cluster creation will fail.
	*/
	NodePoolConfig Dataproc_ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTargetNodePoolConfig `json:"nodePoolConfig,omitempty" yaml:"nodePoolConfig,omitempty"`
}
