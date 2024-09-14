package pubsub

import types "libds/gcp/types"

type LiteSubscription struct {
	/*
	   The settings for this subscription's message delivery.
	   Structure is documented below.
	*/
	DeliveryConfig types.Pubsub_LiteSubscriptionDeliveryConfig `json:"deliveryConfig,omitempty" yaml:"deliveryConfig,omitempty"`

	/*
	   Name of the subscription.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// The region of the pubsub lite topic.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// A reference to a Topic resource.
	Topic string `json:"topic,omitempty" yaml:"topic,omitempty"`

	// The zone of the pubsub lite topic.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}
