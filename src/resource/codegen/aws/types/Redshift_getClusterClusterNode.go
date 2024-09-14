package types

type Redshift_getClusterClusterNode struct {
	// Whether the node is a leader node or a compute node
	NodeRole string `json:"nodeRole,omitempty" yaml:"nodeRole,omitempty"`

	// Private IP address of a node within a cluster
	PrivateIpAddress string `json:"privateIpAddress,omitempty" yaml:"privateIpAddress,omitempty"`

	// Public IP address of a node within a cluster
	PublicIpAddress string `json:"publicIpAddress,omitempty" yaml:"publicIpAddress,omitempty"`
}
