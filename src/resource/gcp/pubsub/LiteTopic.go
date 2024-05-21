package pubsub

import types "DesignSphere_Server/src/resource/gcp/types"

type LiteTopic struct {
	/*
	   Name of the topic.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The settings for this topic's partitions.
	   Structure is documented below.
	*/
	PartitionConfig types.Pubsub_LiteTopicPartitionConfig `json:"partitionConfig,omitempty" yaml:"partitionConfig,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The region of the pubsub lite topic.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	/*
	   The settings for this topic's Reservation usage.
	   Structure is documented below.
	*/
	ReservationConfig types.Pubsub_LiteTopicReservationConfig `json:"reservationConfig,omitempty" yaml:"reservationConfig,omitempty"`

	/*
	   The settings for a topic's message retention.
	   Structure is documented below.
	*/
	RetentionConfig types.Pubsub_LiteTopicRetentionConfig `json:"retentionConfig,omitempty" yaml:"retentionConfig,omitempty"`

	// The zone of the pubsub lite topic.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}
