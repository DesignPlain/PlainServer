package types

type Msk_ReplicatorReplicationInfoListTopicReplication struct {
	// Whether to periodically configure remote topic ACLs to match their corresponding upstream topics.
	CopyAccessControlListsForTopics bool `json:"copyAccessControlListsForTopics,omitempty" yaml:"copyAccessControlListsForTopics,omitempty"`

	// Whether to periodically configure remote topics to match their corresponding upstream topics.
	CopyTopicConfigurations bool `json:"copyTopicConfigurations,omitempty" yaml:"copyTopicConfigurations,omitempty"`

	// Whether to periodically check for new topics and partitions.
	DetectAndCopyNewTopics bool `json:"detectAndCopyNewTopics,omitempty" yaml:"detectAndCopyNewTopics,omitempty"`

	// List of regular expression patterns indicating the topics that should not be replica.
	TopicsToExcludes []string `json:"topicsToExcludes,omitempty" yaml:"topicsToExcludes,omitempty"`

	// List of regular expression patterns indicating the topics to copy.
	TopicsToReplicates []string `json:"topicsToReplicates,omitempty" yaml:"topicsToReplicates,omitempty"`
}
