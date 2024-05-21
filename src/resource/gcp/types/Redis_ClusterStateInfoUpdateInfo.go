package types

type Redis_ClusterStateInfoUpdateInfo struct {
	// Target number of replica nodes per shard.
	TargetReplicaCount int `json:"targetReplicaCount,omitempty" yaml:"targetReplicaCount,omitempty"`

	// Target number of shards for redis cluster.
	TargetShardCount int `json:"targetShardCount,omitempty" yaml:"targetShardCount,omitempty"`
}
