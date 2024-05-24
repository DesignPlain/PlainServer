package types

type Redshift_ScheduledActionTargetActionResizeCluster struct {
	// A boolean value indicating whether the resize operation is using the classic resize process. Default: `false`.
	Classic bool `json:"classic,omitempty" yaml:"classic,omitempty"`

	// The unique identifier for the cluster to resize.
	ClusterIdentifier string `json:"clusterIdentifier,omitempty" yaml:"clusterIdentifier,omitempty"`

	// The new cluster type for the specified cluster.
	ClusterType string `json:"clusterType,omitempty" yaml:"clusterType,omitempty"`

	// The new node type for the nodes you are adding.
	NodeType string `json:"nodeType,omitempty" yaml:"nodeType,omitempty"`

	// The new number of nodes for the cluster.
	NumberOfNodes int `json:"numberOfNodes,omitempty" yaml:"numberOfNodes,omitempty"`
}
