package types

type Gkeonprem_BareMetalClusterNodeConfig struct {
	/*
	   The available runtimes that can be used to run containers in a Bare Metal User Cluster.
	   Possible values are: `CONTAINER_RUNTIME_UNSPECIFIED`, `DOCKER`, `CONTAINERD`.
	*/
	ContainerRuntime string `json:"containerRuntime,omitempty" yaml:"containerRuntime,omitempty"`

	/*
	   The maximum number of pods a node can run. The size of the CIDR range
	   assigned to the node will be derived from this parameter.
	*/
	MaxPodsPerNode int `json:"maxPodsPerNode,omitempty" yaml:"maxPodsPerNode,omitempty"`
}
