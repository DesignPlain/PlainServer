package types

type Pubsub_LiteTopicPartitionConfigCapacity struct {
	// Subscribe throughput capacity per partition in MiB/s. Must be >= 4 and <= 16.
	PublishMibPerSec int `json:"publishMibPerSec,omitempty" yaml:"publishMibPerSec,omitempty"`

	// Publish throughput capacity per partition in MiB/s. Must be >= 4 and <= 16.
	SubscribeMibPerSec int `json:"subscribeMibPerSec,omitempty" yaml:"subscribeMibPerSec,omitempty"`
}
