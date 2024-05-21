package types

type Dataproc_WorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget struct {
	// A namespace within the GKE cluster to deploy into.
	ClusterNamespace string `json:"clusterNamespace,omitempty" yaml:"clusterNamespace,omitempty"`

	// The target GKE cluster to deploy to. Format: 'projects/{project}/locations/{location}/clusters/{cluster_id}'
	TargetGkeCluster string `json:"targetGkeCluster,omitempty" yaml:"targetGkeCluster,omitempty"`
}
