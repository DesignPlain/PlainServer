package types

type Neptune_GlobalClusterGlobalClusterMember struct {
	// Amazon Resource Name (ARN) of member DB Cluster.
	DbClusterArn string `json:"dbClusterArn,omitempty" yaml:"dbClusterArn,omitempty"`

	// Whether the member is the primary DB Cluster.
	IsWriter bool `json:"isWriter,omitempty" yaml:"isWriter,omitempty"`
}
