package types

type Gkeonprem_BareMetalAdminClusterNodeConfig struct {
	/*
	   The maximum number of pods a node can run. The size of the CIDR range
	   assigned to the node will be derived from this parameter.
	*/
	MaxPodsPerNode int `json:"maxPodsPerNode,omitempty" yaml:"maxPodsPerNode,omitempty"`
}
