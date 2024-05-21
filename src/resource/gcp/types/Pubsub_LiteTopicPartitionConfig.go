package types

type Pubsub_LiteTopicPartitionConfig struct {
	/*
	   The capacity configuration.
	   Structure is documented below.
	*/
	Capacity Pubsub_LiteTopicPartitionConfigCapacity `json:"capacity,omitempty" yaml:"capacity,omitempty"`

	// The number of partitions in the topic. Must be at least 1.
	Count int `json:"count,omitempty" yaml:"count,omitempty"`
}
