package types

type Msk_ReplicatorReplicationInfoListConsumerGroupReplication struct {
	// Whether to periodically write the translated offsets to __consumer_offsets topic in target cluster.
	SynchroniseConsumerGroupOffsets bool `json:"synchroniseConsumerGroupOffsets,omitempty" yaml:"synchroniseConsumerGroupOffsets,omitempty"`

	// List of regular expression patterns indicating the consumer groups that should not be replicated.
	ConsumerGroupsToExcludes []string `json:"consumerGroupsToExcludes,omitempty" yaml:"consumerGroupsToExcludes,omitempty"`

	// List of regular expression patterns indicating the consumer groups to copy.
	ConsumerGroupsToReplicates []string `json:"consumerGroupsToReplicates,omitempty" yaml:"consumerGroupsToReplicates,omitempty"`

	// Whether to periodically check for new consumer groups.
	DetectAndCopyNewConsumerGroups bool `json:"detectAndCopyNewConsumerGroups,omitempty" yaml:"detectAndCopyNewConsumerGroups,omitempty"`
}
