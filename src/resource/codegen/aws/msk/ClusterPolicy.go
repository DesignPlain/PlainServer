package msk

type ClusterPolicy struct {
	// The Amazon Resource Name (ARN) that uniquely identifies the cluster.
	ClusterArn string `json:"clusterArn,omitempty" yaml:"clusterArn,omitempty"`

	// Resource policy for cluster.
	Policy string `json:"policy,omitempty" yaml:"policy,omitempty"`
}
