package types

type Elasticache_GlobalReplicationGroupGlobalNodeGroup struct {
	// The ID of the global node group.
	GlobalNodeGroupId string `json:"globalNodeGroupId,omitempty" yaml:"globalNodeGroupId,omitempty"`

	// The keyspace for this node group.
	Slots string `json:"slots,omitempty" yaml:"slots,omitempty"`
}
