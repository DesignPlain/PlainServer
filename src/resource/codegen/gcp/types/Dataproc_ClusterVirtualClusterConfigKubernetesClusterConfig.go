package types

type Dataproc_ClusterVirtualClusterConfigKubernetesClusterConfig struct {
	// The configuration for running the Dataproc cluster on GKE.
	GkeClusterConfig Dataproc_ClusterVirtualClusterConfigKubernetesClusterConfigGkeClusterConfig `json:"gkeClusterConfig,omitempty" yaml:"gkeClusterConfig,omitempty"`

	/*
	   A namespace within the Kubernetes cluster to deploy into.
	   If this namespace does not exist, it is created.
	   If it  exists, Dataproc verifies that another Dataproc VirtualCluster is not installed into it.
	   If not specified, the name of the Dataproc Cluster is used.
	*/
	KubernetesNamespace string `json:"kubernetesNamespace,omitempty" yaml:"kubernetesNamespace,omitempty"`

	// The software configuration for this Dataproc cluster running on Kubernetes.
	KubernetesSoftwareConfig Dataproc_ClusterVirtualClusterConfigKubernetesClusterConfigKubernetesSoftwareConfig `json:"kubernetesSoftwareConfig,omitempty" yaml:"kubernetesSoftwareConfig,omitempty"`
}
