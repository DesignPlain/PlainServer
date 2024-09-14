package types

type Dataproc_ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig struct {
	/*
	   A target GKE cluster to deploy to. It must be in the same project and region as the Dataproc cluster
	   (the GKE cluster can be zonal or regional)
	*/
	GkeClusterTarget string `json:"gkeClusterTarget,omitempty" yaml:"gkeClusterTarget,omitempty"`

	/*
	   GKE node pools where workloads will be scheduled. At least one node pool must be assigned the `DEFAULT`
	   GkeNodePoolTarget.Role. If a GkeNodePoolTarget is not specified, Dataproc constructs a `DEFAULT` GkeNodePoolTarget.
	   Each role can be given to only one GkeNodePoolTarget. All node pools must have the same location settings.
	*/
	NodePoolTargets []Dataproc_ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfigNodePoolTarget `json:"nodePoolTargets,omitempty" yaml:"nodePoolTargets,omitempty"`
}
